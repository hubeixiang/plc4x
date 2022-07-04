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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataNormalRequestGetStatus is the corresponding interface of CALDataNormalRequestGetStatus
type CALDataNormalRequestGetStatus interface {
	utils.LengthAware
	utils.Serializable
	CALDataNormal
	// GetParamNo returns ParamNo (property field)
	GetParamNo() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
}

// CALDataNormalRequestGetStatusExactly can be used when we want exactly this type and not a type which fulfills CALDataNormalRequestGetStatus.
// This is useful for switch cases.
type CALDataNormalRequestGetStatusExactly interface {
	CALDataNormalRequestGetStatus
	isCALDataNormalRequestGetStatus() bool
}

// _CALDataNormalRequestGetStatus is the data-structure of this message
type _CALDataNormalRequestGetStatus struct {
	*_CALDataNormal
	ParamNo uint8
	Count   uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataNormalRequestGetStatus) InitializeParent(parent CALDataNormal, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataNormalRequestGetStatus) GetParent() CALDataNormal {
	return m._CALDataNormal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataNormalRequestGetStatus) GetParamNo() uint8 {
	return m.ParamNo
}

func (m *_CALDataNormalRequestGetStatus) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataNormalRequestGetStatus factory function for _CALDataNormalRequestGetStatus
func NewCALDataNormalRequestGetStatus(paramNo uint8, count uint8, commandTypeContainer CALCommandTypeContainer) *_CALDataNormalRequestGetStatus {
	_result := &_CALDataNormalRequestGetStatus{
		ParamNo:        paramNo,
		Count:          count,
		_CALDataNormal: NewCALDataNormal(commandTypeContainer),
	}
	_result._CALDataNormal._CALDataNormalChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataNormalRequestGetStatus(structType interface{}) CALDataNormalRequestGetStatus {
	if casted, ok := structType.(CALDataNormalRequestGetStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataNormalRequestGetStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataNormalRequestGetStatus) GetTypeName() string {
	return "CALDataNormalRequestGetStatus"
}

func (m *_CALDataNormalRequestGetStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataNormalRequestGetStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataNormalRequestGetStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataNormalRequestGetStatusParse(readBuffer utils.ReadBuffer) (CALDataNormalRequestGetStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataNormalRequestGetStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataNormalRequestGetStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	_paramNo, _paramNoErr := readBuffer.ReadUint8("paramNo", 8)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataNormalRequestGetStatus")
	}
	paramNo := _paramNo

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of CALDataNormalRequestGetStatus")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("CALDataNormalRequestGetStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataNormalRequestGetStatus")
	}

	// Create a partially initialized instance
	_child := &_CALDataNormalRequestGetStatus{
		ParamNo:        paramNo,
		Count:          count,
		_CALDataNormal: &_CALDataNormal{},
	}
	_child._CALDataNormal._CALDataNormalChildRequirements = _child
	return _child, nil
}

func (m *_CALDataNormalRequestGetStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataNormalRequestGetStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataNormalRequestGetStatus")
		}

		// Simple Field (paramNo)
		paramNo := uint8(m.GetParamNo())
		_paramNoErr := writeBuffer.WriteUint8("paramNo", 8, (paramNo))
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := writeBuffer.WriteUint8("count", 8, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataNormalRequestGetStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataNormalRequestGetStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataNormalRequestGetStatus) isCALDataNormalRequestGetStatus() bool {
	return true
}

func (m *_CALDataNormalRequestGetStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}