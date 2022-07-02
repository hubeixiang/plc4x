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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestCommand_INITIATOR byte = 0x5C

// RequestCommand is the corresponding interface of RequestCommand
type RequestCommand interface {
	utils.LengthAware
	utils.Serializable
	Request
	// GetCbusCommand returns CbusCommand (property field)
	GetCbusCommand() CBusCommand
}

// RequestCommandExactly can be used when we want exactly this type and not a type which fulfills RequestCommand.
// This is useful for switch cases.
type RequestCommandExactly interface {
	RequestCommand
	isRequestCommand() bool
}

// _RequestCommand is the data-structure of this message
type _RequestCommand struct {
	*_Request
	CbusCommand CBusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestCommand) InitializeParent(parent Request, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_RequestCommand) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestCommand) GetCbusCommand() CBusCommand {
	return m.CbusCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestCommand) GetInitiator() byte {
	return RequestCommand_INITIATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestCommand factory function for _RequestCommand
func NewRequestCommand(cbusCommand CBusCommand, peekedByte byte, srchk bool) *_RequestCommand {
	_result := &_RequestCommand{
		CbusCommand: cbusCommand,
		_Request:    NewRequest(peekedByte, srchk),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestCommand(structType interface{}) RequestCommand {
	if casted, ok := structType.(RequestCommand); ok {
		return casted
	}
	if casted, ok := structType.(*RequestCommand); ok {
		return *casted
	}
	return nil
}

func (m *_RequestCommand) GetTypeName() string {
	return "RequestCommand"
}

func (m *_RequestCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (initiator)
	lengthInBits += 8

	// Simple field (cbusCommand)
	lengthInBits += m.CbusCommand.GetLengthInBits()

	return lengthInBits
}

func (m *_RequestCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestCommandParse(readBuffer utils.ReadBuffer, srchk bool) (RequestCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (initiator)
	initiator, _initiatorErr := readBuffer.ReadByte("initiator")
	if _initiatorErr != nil {
		return nil, errors.Wrap(_initiatorErr, "Error parsing 'initiator' field of RequestCommand")
	}
	if initiator != RequestCommand_INITIATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestCommand_INITIATOR) + " but got " + fmt.Sprintf("%d", initiator))
	}

	// Simple Field (cbusCommand)
	if pullErr := readBuffer.PullContext("cbusCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for cbusCommand")
	}
	_cbusCommand, _cbusCommandErr := CBusCommandParse(readBuffer, bool(srchk))
	if _cbusCommandErr != nil {
		return nil, errors.Wrap(_cbusCommandErr, "Error parsing 'cbusCommand' field of RequestCommand")
	}
	cbusCommand := _cbusCommand.(CBusCommand)
	if closeErr := readBuffer.CloseContext("cbusCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for cbusCommand")
	}

	if closeErr := readBuffer.CloseContext("RequestCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestCommand")
	}

	// Create a partially initialized instance
	_child := &_RequestCommand{
		CbusCommand: cbusCommand,
		_Request: &_Request{
			Srchk: srchk,
		},
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestCommand")
		}

		// Const Field (initiator)
		_initiatorErr := writeBuffer.WriteByte("initiator", 0x5C)
		if _initiatorErr != nil {
			return errors.Wrap(_initiatorErr, "Error serializing 'initiator' field")
		}

		// Simple Field (cbusCommand)
		if pushErr := writeBuffer.PushContext("cbusCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for cbusCommand")
		}
		_cbusCommandErr := writeBuffer.WriteSerializable(m.GetCbusCommand())
		if popErr := writeBuffer.PopContext("cbusCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for cbusCommand")
		}
		if _cbusCommandErr != nil {
			return errors.Wrap(_cbusCommandErr, "Error serializing 'cbusCommand' field")
		}

		if popErr := writeBuffer.PopContext("RequestCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestCommand) isRequestCommand() bool {
	return true
}

func (m *_RequestCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}