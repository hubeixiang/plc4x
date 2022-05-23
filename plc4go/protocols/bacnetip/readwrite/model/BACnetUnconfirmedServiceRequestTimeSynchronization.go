/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestTimeSynchronization is the data-structure of this message
type BACnetUnconfirmedServiceRequestTimeSynchronization struct {
	*BACnetUnconfirmedServiceRequest
	SynchronizedDate *BACnetApplicationTagDate
	SynchronizedTime *BACnetApplicationTagTime

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestTimeSynchronization
type IBACnetUnconfirmedServiceRequestTimeSynchronization interface {
	IBACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() *BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() *BACnetApplicationTagTime
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedDate() *BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedTime() *BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestTimeSynchronization factory function for BACnetUnconfirmedServiceRequestTimeSynchronization
func NewBACnetUnconfirmedServiceRequestTimeSynchronization(synchronizedDate *BACnetApplicationTagDate, synchronizedTime *BACnetApplicationTagTime, serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestTimeSynchronization {
	_result := &BACnetUnconfirmedServiceRequestTimeSynchronization{
		SynchronizedDate:                synchronizedDate,
		SynchronizedTime:                synchronizedTime,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestTimeSynchronization(structType interface{}) *BACnetUnconfirmedServiceRequestTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestTimeSynchronization(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestTimeSynchronization(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestTimeSynchronization"
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits()

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestTimeSynchronizationParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestTimeSynchronization, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (synchronizedDate)
	if pullErr := readBuffer.PullContext("synchronizedDate"); pullErr != nil {
		return nil, pullErr
	}
	_synchronizedDate, _synchronizedDateErr := BACnetApplicationTagParse(readBuffer)
	if _synchronizedDateErr != nil {
		return nil, errors.Wrap(_synchronizedDateErr, "Error parsing 'synchronizedDate' field")
	}
	synchronizedDate := CastBACnetApplicationTagDate(_synchronizedDate)
	if closeErr := readBuffer.CloseContext("synchronizedDate"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (synchronizedTime)
	if pullErr := readBuffer.PullContext("synchronizedTime"); pullErr != nil {
		return nil, pullErr
	}
	_synchronizedTime, _synchronizedTimeErr := BACnetApplicationTagParse(readBuffer)
	if _synchronizedTimeErr != nil {
		return nil, errors.Wrap(_synchronizedTimeErr, "Error parsing 'synchronizedTime' field")
	}
	synchronizedTime := CastBACnetApplicationTagTime(_synchronizedTime)
	if closeErr := readBuffer.CloseContext("synchronizedTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestTimeSynchronization{
		SynchronizedDate:                CastBACnetApplicationTagDate(synchronizedDate),
		SynchronizedTime:                CastBACnetApplicationTagTime(synchronizedTime),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pushErr != nil {
			return pushErr
		}

		// Simple Field (synchronizedDate)
		if pushErr := writeBuffer.PushContext("synchronizedDate"); pushErr != nil {
			return pushErr
		}
		_synchronizedDateErr := m.SynchronizedDate.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedDate"); popErr != nil {
			return popErr
		}
		if _synchronizedDateErr != nil {
			return errors.Wrap(_synchronizedDateErr, "Error serializing 'synchronizedDate' field")
		}

		// Simple Field (synchronizedTime)
		if pushErr := writeBuffer.PushContext("synchronizedTime"); pushErr != nil {
			return pushErr
		}
		_synchronizedTimeErr := m.SynchronizedTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedTime"); popErr != nil {
			return popErr
		}
		if _synchronizedTimeErr != nil {
			return errors.Wrap(_synchronizedTimeErr, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestTimeSynchronization) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}