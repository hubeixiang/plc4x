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
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataIndividualAddressRead is the data-structure of this message
type ApduDataIndividualAddressRead struct {
	*ApduData

	// Arguments.
	DataLength uint8
}

// IApduDataIndividualAddressRead is the corresponding interface of ApduDataIndividualAddressRead
type IApduDataIndividualAddressRead interface {
	IApduData
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

func (m *ApduDataIndividualAddressRead) GetApciType() uint8 {
	return 0x4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataIndividualAddressRead) InitializeParent(parent *ApduData) {}

func (m *ApduDataIndividualAddressRead) GetParent() *ApduData {
	return m.ApduData
}

// NewApduDataIndividualAddressRead factory function for ApduDataIndividualAddressRead
func NewApduDataIndividualAddressRead(dataLength uint8) *ApduDataIndividualAddressRead {
	_result := &ApduDataIndividualAddressRead{
		ApduData: NewApduData(dataLength),
	}
	_result.Child = _result
	return _result
}

func CastApduDataIndividualAddressRead(structType interface{}) *ApduDataIndividualAddressRead {
	if casted, ok := structType.(ApduDataIndividualAddressRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataIndividualAddressRead(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataIndividualAddressRead(casted.Child)
	}
	return nil
}

func (m *ApduDataIndividualAddressRead) GetTypeName() string {
	return "ApduDataIndividualAddressRead"
}

func (m *ApduDataIndividualAddressRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataIndividualAddressRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataIndividualAddressRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataIndividualAddressReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduDataIndividualAddressRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataIndividualAddressRead{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child, nil
}

func (m *ApduDataIndividualAddressRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataIndividualAddressRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}