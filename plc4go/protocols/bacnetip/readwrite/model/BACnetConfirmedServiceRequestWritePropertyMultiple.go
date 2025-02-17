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

// BACnetConfirmedServiceRequestWritePropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestWritePropertyMultiple
type BACnetConfirmedServiceRequestWritePropertyMultiple interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []BACnetWriteAccessSpecification
}

// BACnetConfirmedServiceRequestWritePropertyMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestWritePropertyMultiple.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestWritePropertyMultipleExactly interface {
	BACnetConfirmedServiceRequestWritePropertyMultiple
	isBACnetConfirmedServiceRequestWritePropertyMultiple() bool
}

// _BACnetConfirmedServiceRequestWritePropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestWritePropertyMultiple struct {
	*_BACnetConfirmedServiceRequest
	Data []BACnetWriteAccessSpecification

	// Arguments.
	ServiceRequestPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetData() []BACnetWriteAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestWritePropertyMultiple factory function for _BACnetConfirmedServiceRequestWritePropertyMultiple
func NewBACnetConfirmedServiceRequestWritePropertyMultiple(data []BACnetWriteAccessSpecification, serviceRequestLength uint16, serviceRequestPayloadLength uint16) *_BACnetConfirmedServiceRequestWritePropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		Data:                           data,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWritePropertyMultiple(structType interface{}) BACnetConfirmedServiceRequestWritePropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWritePropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestWritePropertyMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16, serviceRequestPayloadLength uint16) (BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Length array
	var data []BACnetWriteAccessSpecification
	{
		_dataLength := serviceRequestPayloadLength
		_dataEndPos := positionAware.GetPos() + uint16(_dataLength)
		for positionAware.GetPos() < _dataEndPos {
			_item, _err := BACnetWriteAccessSpecificationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of BACnetConfirmedServiceRequestWritePropertyMultiple")
			}
			data = append(data, _item.(BACnetWriteAccessSpecification))
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWritePropertyMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestWritePropertyMultiple{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		Data: data,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWritePropertyMultiple")
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

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWritePropertyMultiple")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceRequestPayloadLength() uint16 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) isBACnetConfirmedServiceRequestWritePropertyMultiple() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestWritePropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
