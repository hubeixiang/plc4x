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

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestReadPropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestReadPropertyMultiple
type BACnetConfirmedServiceRequestReadPropertyMultiple interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []BACnetReadAccessSpecification
}

// BACnetConfirmedServiceRequestReadPropertyMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadPropertyMultiple.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadPropertyMultipleExactly interface {
	BACnetConfirmedServiceRequestReadPropertyMultiple
	isBACnetConfirmedServiceRequestReadPropertyMultiple() bool
}

// _BACnetConfirmedServiceRequestReadPropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestReadPropertyMultiple struct {
	*_BACnetConfirmedServiceRequest
	Data []BACnetReadAccessSpecification

	// Arguments.
	ServiceRequestPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetData() []BACnetReadAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadPropertyMultiple factory function for _BACnetConfirmedServiceRequestReadPropertyMultiple
func NewBACnetConfirmedServiceRequestReadPropertyMultiple(data []BACnetReadAccessSpecification, serviceRequestLength uint16, serviceRequestPayloadLength uint16) *_BACnetConfirmedServiceRequestReadPropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestReadPropertyMultiple{
		Data:                           data,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadPropertyMultiple(structType interface{}) BACnetConfirmedServiceRequestReadPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadPropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadPropertyMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16, serviceRequestPayloadLength uint16) (BACnetConfirmedServiceRequestReadPropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Length array
	var data []BACnetReadAccessSpecification
	{
		_dataLength := serviceRequestPayloadLength
		_dataEndPos := positionAware.GetPos() + uint16(_dataLength)
		for positionAware.GetPos() < _dataEndPos {
			_item, _err := BACnetReadAccessSpecificationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of BACnetConfirmedServiceRequestReadPropertyMultiple")
			}
			data = append(data, _item.(BACnetReadAccessSpecification))
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadPropertyMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadPropertyMultiple{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		Data: data,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadPropertyMultiple")
		}

		// Array Field (data)
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		for _, _element := range m.GetData() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadPropertyMultiple")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) GetServiceRequestPayloadLength() uint16 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) isBACnetConfirmedServiceRequestReadPropertyMultiple() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
