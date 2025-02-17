/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package cbus

import (
	"context"
	"fmt"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
	_default "github.com/apache/plc4x/plc4go/spi/default"
	"github.com/apache/plc4x/plc4go/spi/model"
	"github.com/rs/zerolog/log"
)

type Browser struct {
	_default.DefaultBrowser
	connection      *Connection
	messageCodec    spi.MessageCodec
	sequenceCounter uint8
}

func NewBrowser(connection *Connection, messageCodec spi.MessageCodec) *Browser {
	browser := Browser{
		connection:      connection,
		messageCodec:    messageCodec,
		sequenceCounter: 0,
	}
	browser.DefaultBrowser = _default.NewDefaultBrowser(browser)
	return &browser
}

func (m Browser) BrowseField(ctx context.Context, browseRequest apiModel.PlcBrowseRequest, interceptor func(result apiModel.PlcBrowseEvent) bool, fieldName string, field apiModel.PlcField) (apiModel.PlcResponseCode, []apiModel.PlcBrowseFoundField) {
	// TODO: handle ctx
	var queryResults []apiModel.PlcBrowseFoundField
	switch field := field.(type) {
	case *unitInfoField:
		allUnits := false
		var units []readWriteModel.UnitAddress
		allAttributes := false
		var attributes []readWriteModel.Attribute
		if unitAddress := field.unitAddress; unitAddress != nil {
			units = append(units, *unitAddress)
		} else {
			allUnits = true
			for i := 0; i <= 0xFF; i++ {
				units = append(units, readWriteModel.NewUnitAddress(byte(i)))
			}
		}
		if attribute := field.attribute; attribute != nil {
			attributes = append(attributes, *attribute)
		} else {
			allAttributes = true
			for _, attribute := range readWriteModel.AttributeValues {
				attributes = append(attributes, attribute)
			}
		}

		if allUnits {
			log.Info().Msg("Querying all units")
		}
	unitLoop:
		for _, unit := range units {
			unitAddress := unit.GetAddress()
			if !allUnits && allAttributes {
				log.Info().Msgf("Querying all attributes of unit %d", unitAddress)
			}
			event := log.Info()
			if allUnits {
				event = log.Debug()
			}
			event.Msgf("Query unit  %d", unitAddress)
			for _, attribute := range attributes {
				if !allUnits && !allAttributes {
					log.Info().Msgf("Querying attribute %s of unit %d", attribute, unitAddress)
				} else {
					event.Msgf("unit %d: Query %s", unitAddress, attribute)
				}
				readFieldName := fmt.Sprintf("%s/%d/%s", fieldName, unitAddress, attribute)
				readRequest, _ := m.connection.ReadRequestBuilder().
					AddField(readFieldName, NewCALIdentifyField(unit, attribute, 1)).
					Build()
				requestResult := <-readRequest.Execute()
				if err := requestResult.GetErr(); err != nil {
					if !allUnits && !allAttributes {
						event.Err(err).Msgf("unit %d: Can't read attribute %s", unitAddress, attribute)
					}
					continue unitLoop
				}
				response := requestResult.GetResponse()
				if code := response.GetResponseCode(readFieldName); code != apiModel.PlcResponseCode_OK {
					event.Msgf("unit %d: error reading field %s. Code %s", unitAddress, attribute, code)
					continue unitLoop
				}
				queryResult := &model.DefaultPlcBrowseQueryResult{
					Field:        NewCALIdentifyField(unit, attribute, 1),
					Name:         fieldName,
					Readable:     true,
					Writable:     false,
					Subscribable: false,
					Attributes: map[string]values.PlcValue{
						"CurrentValue": response.GetValue(readFieldName),
					},
				}
				if interceptor != nil {
					interceptor(&model.DefaultPlcBrowseEvent{
						Request:   browseRequest,
						FieldName: readFieldName,
						Result:    queryResult,
						Err:       nil,
					})
				}
				queryResults = append(queryResults, queryResult)
			}
		}
	default:
		return apiModel.PlcResponseCode_INTERNAL_ERROR, nil
	}
	return apiModel.PlcResponseCode_OK, queryResults
}
