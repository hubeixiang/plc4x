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

// BACnetHostAddressIpAddress is the corresponding interface of BACnetHostAddressIpAddress
type BACnetHostAddressIpAddress interface {
	utils.LengthAware
	utils.Serializable
	BACnetHostAddress
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() BACnetContextTagOctetString
}

// BACnetHostAddressIpAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetHostAddressIpAddress.
// This is useful for switch cases.
type BACnetHostAddressIpAddressExactly interface {
	BACnetHostAddressIpAddress
	isBACnetHostAddressIpAddress() bool
}

// _BACnetHostAddressIpAddress is the data-structure of this message
type _BACnetHostAddressIpAddress struct {
	*_BACnetHostAddress
	IpAddress BACnetContextTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetHostAddressIpAddress) InitializeParent(parent BACnetHostAddress, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetHostAddressIpAddress) GetParent() BACnetHostAddress {
	return m._BACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressIpAddress) GetIpAddress() BACnetContextTagOctetString {
	return m.IpAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressIpAddress factory function for _BACnetHostAddressIpAddress
func NewBACnetHostAddressIpAddress(ipAddress BACnetContextTagOctetString, peekedTagHeader BACnetTagHeader) *_BACnetHostAddressIpAddress {
	_result := &_BACnetHostAddressIpAddress{
		IpAddress:          ipAddress,
		_BACnetHostAddress: NewBACnetHostAddress(peekedTagHeader),
	}
	_result._BACnetHostAddress._BACnetHostAddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressIpAddress(structType interface{}) BACnetHostAddressIpAddress {
	if casted, ok := structType.(BACnetHostAddressIpAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressIpAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressIpAddress) GetTypeName() string {
	return "BACnetHostAddressIpAddress"
}

func (m *_BACnetHostAddressIpAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetHostAddressIpAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetHostAddressIpAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostAddressIpAddressParse(readBuffer utils.ReadBuffer) (BACnetHostAddressIpAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressIpAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressIpAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipAddress)
	if pullErr := readBuffer.PullContext("ipAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipAddress")
	}
	_ipAddress, _ipAddressErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_OCTET_STRING))
	if _ipAddressErr != nil {
		return nil, errors.Wrap(_ipAddressErr, "Error parsing 'ipAddress' field of BACnetHostAddressIpAddress")
	}
	ipAddress := _ipAddress.(BACnetContextTagOctetString)
	if closeErr := readBuffer.CloseContext("ipAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipAddress")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddressIpAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressIpAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetHostAddressIpAddress{
		_BACnetHostAddress: &_BACnetHostAddress{},
		IpAddress:          ipAddress,
	}
	_child._BACnetHostAddress._BACnetHostAddressChildRequirements = _child
	return _child, nil
}

func (m *_BACnetHostAddressIpAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressIpAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressIpAddress")
		}

		// Simple Field (ipAddress)
		if pushErr := writeBuffer.PushContext("ipAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipAddress")
		}
		_ipAddressErr := writeBuffer.WriteSerializable(m.GetIpAddress())
		if popErr := writeBuffer.PopContext("ipAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipAddress")
		}
		if _ipAddressErr != nil {
			return errors.Wrap(_ipAddressErr, "Error serializing 'ipAddress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressIpAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressIpAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetHostAddressIpAddress) isBACnetHostAddressIpAddress() bool {
	return true
}

func (m *_BACnetHostAddressIpAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
