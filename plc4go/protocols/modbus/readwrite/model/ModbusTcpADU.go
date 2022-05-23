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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusTcpADU_PROTOCOLIDENTIFIER uint16 = 0x0000

// ModbusTcpADU is the data-structure of this message
type ModbusTcpADU struct {
	*ModbusADU
	TransactionIdentifier uint16
	UnitIdentifier        uint8
	Pdu                   *ModbusPDU

	// Arguments.
	Response bool
}

// IModbusTcpADU is the corresponding interface of ModbusTcpADU
type IModbusTcpADU interface {
	IModbusADU
	// GetTransactionIdentifier returns TransactionIdentifier (property field)
	GetTransactionIdentifier() uint16
	// GetUnitIdentifier returns UnitIdentifier (property field)
	GetUnitIdentifier() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() *ModbusPDU
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

func (m *ModbusTcpADU) GetDriverType() DriverType {
	return DriverType_MODBUS_TCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusTcpADU) InitializeParent(parent *ModbusADU) {}

func (m *ModbusTcpADU) GetParent() *ModbusADU {
	return m.ModbusADU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusTcpADU) GetTransactionIdentifier() uint16 {
	return m.TransactionIdentifier
}

func (m *ModbusTcpADU) GetUnitIdentifier() uint8 {
	return m.UnitIdentifier
}

func (m *ModbusTcpADU) GetPdu() *ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *ModbusTcpADU) GetProtocolIdentifier() uint16 {
	return ModbusTcpADU_PROTOCOLIDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusTcpADU factory function for ModbusTcpADU
func NewModbusTcpADU(transactionIdentifier uint16, unitIdentifier uint8, pdu *ModbusPDU, response bool) *ModbusTcpADU {
	_result := &ModbusTcpADU{
		TransactionIdentifier: transactionIdentifier,
		UnitIdentifier:        unitIdentifier,
		Pdu:                   pdu,
		ModbusADU:             NewModbusADU(response),
	}
	_result.Child = _result
	return _result
}

func CastModbusTcpADU(structType interface{}) *ModbusTcpADU {
	if casted, ok := structType.(ModbusTcpADU); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusTcpADU); ok {
		return casted
	}
	if casted, ok := structType.(ModbusADU); ok {
		return CastModbusTcpADU(casted.Child)
	}
	if casted, ok := structType.(*ModbusADU); ok {
		return CastModbusTcpADU(casted.Child)
	}
	return nil
}

func (m *ModbusTcpADU) GetTypeName() string {
	return "ModbusTcpADU"
}

func (m *ModbusTcpADU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusTcpADU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (transactionIdentifier)
	lengthInBits += 16

	// Const Field (protocolIdentifier)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 16

	// Simple field (unitIdentifier)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits()

	return lengthInBits
}

func (m *ModbusTcpADU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusTcpADUParse(readBuffer utils.ReadBuffer, driverType DriverType, response bool) (*ModbusTcpADU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusTcpADU"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (transactionIdentifier)
	_transactionIdentifier, _transactionIdentifierErr := readBuffer.ReadUint16("transactionIdentifier", 16)
	if _transactionIdentifierErr != nil {
		return nil, errors.Wrap(_transactionIdentifierErr, "Error parsing 'transactionIdentifier' field")
	}
	transactionIdentifier := _transactionIdentifier

	// Const Field (protocolIdentifier)
	protocolIdentifier, _protocolIdentifierErr := readBuffer.ReadUint16("protocolIdentifier", 16)
	if _protocolIdentifierErr != nil {
		return nil, errors.Wrap(_protocolIdentifierErr, "Error parsing 'protocolIdentifier' field")
	}
	if protocolIdentifier != ModbusTcpADU_PROTOCOLIDENTIFIER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ModbusTcpADU_PROTOCOLIDENTIFIER) + " but got " + fmt.Sprintf("%d", protocolIdentifier))
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint16("length", 16)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (unitIdentifier)
	_unitIdentifier, _unitIdentifierErr := readBuffer.ReadUint8("unitIdentifier", 8)
	if _unitIdentifierErr != nil {
		return nil, errors.Wrap(_unitIdentifierErr, "Error parsing 'unitIdentifier' field")
	}
	unitIdentifier := _unitIdentifier

	// Simple Field (pdu)
	if pullErr := readBuffer.PullContext("pdu"); pullErr != nil {
		return nil, pullErr
	}
	_pdu, _pduErr := ModbusPDUParse(readBuffer, bool(response))
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field")
	}
	pdu := CastModbusPDU(_pdu)
	if closeErr := readBuffer.CloseContext("pdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusTcpADU"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusTcpADU{
		TransactionIdentifier: transactionIdentifier,
		UnitIdentifier:        unitIdentifier,
		Pdu:                   CastModbusPDU(pdu),
		ModbusADU:             &ModbusADU{},
	}
	_child.ModbusADU.Child = _child
	return _child, nil
}

func (m *ModbusTcpADU) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusTcpADU"); pushErr != nil {
			return pushErr
		}

		// Simple Field (transactionIdentifier)
		transactionIdentifier := uint16(m.TransactionIdentifier)
		_transactionIdentifierErr := writeBuffer.WriteUint16("transactionIdentifier", 16, (transactionIdentifier))
		if _transactionIdentifierErr != nil {
			return errors.Wrap(_transactionIdentifierErr, "Error serializing 'transactionIdentifier' field")
		}

		// Const Field (protocolIdentifier)
		_protocolIdentifierErr := writeBuffer.WriteUint16("protocolIdentifier", 16, 0x0000)
		if _protocolIdentifierErr != nil {
			return errors.Wrap(_protocolIdentifierErr, "Error serializing 'protocolIdentifier' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint16(uint16(m.GetPdu().GetLengthInBytes()) + uint16(uint16(1)))
		_lengthErr := writeBuffer.WriteUint16("length", 16, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (unitIdentifier)
		unitIdentifier := uint8(m.UnitIdentifier)
		_unitIdentifierErr := writeBuffer.WriteUint8("unitIdentifier", 8, (unitIdentifier))
		if _unitIdentifierErr != nil {
			return errors.Wrap(_unitIdentifierErr, "Error serializing 'unitIdentifier' field")
		}

		// Simple Field (pdu)
		if pushErr := writeBuffer.PushContext("pdu"); pushErr != nil {
			return pushErr
		}
		_pduErr := m.Pdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("pdu"); popErr != nil {
			return popErr
		}
		if _pduErr != nil {
			return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
		}

		if popErr := writeBuffer.PopContext("ModbusTcpADU"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusTcpADU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}