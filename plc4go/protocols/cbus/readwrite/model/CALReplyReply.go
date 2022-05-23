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

// CALReplyReply is the data-structure of this message
type CALReplyReply struct {
	*Reply
	IsA *CALReply
}

// ICALReplyReply is the corresponding interface of CALReplyReply
type ICALReplyReply interface {
	IReply
	// GetIsA returns IsA (property field)
	GetIsA() *CALReply
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CALReplyReply) InitializeParent(parent *Reply, magicByte byte) {
	m.Reply.MagicByte = magicByte
}

func (m *CALReplyReply) GetParent() *Reply {
	return m.Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALReplyReply) GetIsA() *CALReply {
	return m.IsA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReplyReply factory function for CALReplyReply
func NewCALReplyReply(isA *CALReply, magicByte byte) *CALReplyReply {
	_result := &CALReplyReply{
		IsA:   isA,
		Reply: NewReply(magicByte),
	}
	_result.Child = _result
	return _result
}

func CastCALReplyReply(structType interface{}) *CALReplyReply {
	if casted, ok := structType.(CALReplyReply); ok {
		return &casted
	}
	if casted, ok := structType.(*CALReplyReply); ok {
		return casted
	}
	if casted, ok := structType.(Reply); ok {
		return CastCALReplyReply(casted.Child)
	}
	if casted, ok := structType.(*Reply); ok {
		return CastCALReplyReply(casted.Child)
	}
	return nil
}

func (m *CALReplyReply) GetTypeName() string {
	return "CALReplyReply"
}

func (m *CALReplyReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALReplyReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isA)
	lengthInBits += m.IsA.GetLengthInBits()

	return lengthInBits
}

func (m *CALReplyReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyReplyParse(readBuffer utils.ReadBuffer) (*CALReplyReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyReply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isA)
	if pullErr := readBuffer.PullContext("isA"); pullErr != nil {
		return nil, pullErr
	}
	_isA, _isAErr := CALReplyParse(readBuffer)
	if _isAErr != nil {
		return nil, errors.Wrap(_isAErr, "Error parsing 'isA' field")
	}
	isA := CastCALReply(_isA)
	if closeErr := readBuffer.CloseContext("isA"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CALReplyReply"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALReplyReply{
		IsA:   CastCALReply(isA),
		Reply: &Reply{},
	}
	_child.Reply.Child = _child
	return _child, nil
}

func (m *CALReplyReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyReply"); pushErr != nil {
			return pushErr
		}

		// Simple Field (isA)
		if pushErr := writeBuffer.PushContext("isA"); pushErr != nil {
			return pushErr
		}
		_isAErr := m.IsA.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("isA"); popErr != nil {
			return popErr
		}
		if _isAErr != nil {
			return errors.Wrap(_isAErr, "Error serializing 'isA' field")
		}

		if popErr := writeBuffer.PopContext("CALReplyReply"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALReplyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}