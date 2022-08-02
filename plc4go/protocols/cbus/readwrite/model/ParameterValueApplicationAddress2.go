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

// ParameterValueApplicationAddress2 is the corresponding interface of ParameterValueApplicationAddress2
type ParameterValueApplicationAddress2 interface {
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() ApplicationAddress1
}

// ParameterValueApplicationAddress2Exactly can be used when we want exactly this type and not a type which fulfills ParameterValueApplicationAddress2.
// This is useful for switch cases.
type ParameterValueApplicationAddress2Exactly interface {
	ParameterValueApplicationAddress2
	isParameterValueApplicationAddress2() bool
}

// _ParameterValueApplicationAddress2 is the data-structure of this message
type _ParameterValueApplicationAddress2 struct {
	*_ParameterValue
	Value ApplicationAddress1
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetParameterType() ParameterType {
	return ParameterType_APPLICATION_ADDRESS_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueApplicationAddress2) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueApplicationAddress2) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetValue() ApplicationAddress1 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueApplicationAddress2 factory function for _ParameterValueApplicationAddress2
func NewParameterValueApplicationAddress2(value ApplicationAddress1, numBytes uint8) *_ParameterValueApplicationAddress2 {
	_result := &_ParameterValueApplicationAddress2{
		Value:           value,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueApplicationAddress2(structType interface{}) ParameterValueApplicationAddress2 {
	if casted, ok := structType.(ParameterValueApplicationAddress2); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueApplicationAddress2); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueApplicationAddress2) GetTypeName() string {
	return "ParameterValueApplicationAddress2"
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	return lengthInBits
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterValueApplicationAddress2Parse(readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueApplicationAddress2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueApplicationAddress2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueApplicationAddress2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) == (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{"ApplicationAddress2 has exactly one byte"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := ApplicationAddress1Parse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueApplicationAddress2")
	}
	value := _value.(ApplicationAddress1)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueApplicationAddress2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueApplicationAddress2")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueApplicationAddress2{
		Value: value,
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueApplicationAddress2) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueApplicationAddress2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueApplicationAddress2")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueApplicationAddress2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueApplicationAddress2")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ParameterValueApplicationAddress2) isParameterValueApplicationAddress2() bool {
	return true
}

func (m *_ParameterValueApplicationAddress2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}