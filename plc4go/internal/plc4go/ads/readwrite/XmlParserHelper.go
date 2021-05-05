//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package readwrite

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

type AdsXmlParserHelper struct {
}

func (m AdsXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	_ = strconv.Atoi
	switch typeName {
	case "AdsMultiRequestItem":
		return nil, errors.New("AdsMultiRequestItem unmappable")
	case "AmsTCPPacket":
		return model.AmsTCPPacketParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "State":
		return model.StateParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsPacket":
		return model.AmsPacketParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsSerialFrame":
		return model.AmsSerialFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsSerialAcknowledgeFrame":
		return model.AmsSerialAcknowledgeFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsData":
		return nil, errors.New("AdsData unmappable")
	case "AmsNetId":
		return model.AmsNetIdParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsStampHeader":
		return model.AdsStampHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AmsSerialResetFrame":
		return model.AmsSerialResetFrameParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AdsNotificationSample":
		return model.AdsNotificationSampleParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}

// Deprecated: will be removed in favor of Parse soon
func (m AdsXmlParserHelper) ParseOld(typeName string, xmlString string) (interface{}, error) {
	switch typeName {
	case "AdsMultiRequestItem":
		var obj *model.AdsMultiRequestItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsTCPPacket":
		var obj *model.AmsTCPPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "State":
		var obj *model.State
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsPacket":
		var obj *model.AmsPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsSerialFrame":
		var obj *model.AmsSerialFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsSerialAcknowledgeFrame":
		var obj *model.AmsSerialAcknowledgeFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AdsData":
		var obj *model.AdsData
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsNetId":
		var obj *model.AmsNetId
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AdsStampHeader":
		var obj *model.AdsStampHeader
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AmsSerialResetFrame":
		var obj *model.AmsSerialResetFrame
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "AdsNotificationSample":
		var obj *model.AdsNotificationSample
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}