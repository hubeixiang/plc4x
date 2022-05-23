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

// BACnetNotificationParametersChangeOfValueNewValue is the data-structure of this message
type BACnetNotificationParametersChangeOfValueNewValue struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
	Child     IBACnetNotificationParametersChangeOfValueNewValueChild
}

// IBACnetNotificationParametersChangeOfValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValue
type IBACnetNotificationParametersChangeOfValueNewValue interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetNotificationParametersChangeOfValueNewValueParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParametersChangeOfValueNewValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetNotificationParametersChangeOfValueNewValueChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetNotificationParametersChangeOfValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag)
	GetParent() *BACnetNotificationParametersChangeOfValueNewValue

	GetTypeName() string
	IBACnetNotificationParametersChangeOfValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValue factory function for BACnetNotificationParametersChangeOfValueNewValue
func NewBACnetNotificationParametersChangeOfValueNewValue(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfValueNewValue {
	return &BACnetNotificationParametersChangeOfValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetNotificationParametersChangeOfValueNewValue(structType interface{}) *BACnetNotificationParametersChangeOfValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetNotificationParametersChangeOfValueNewValueChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValue"
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfValueNewValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfValueNewValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetNotificationParametersChangeOfValueNewValueChild interface {
		InitializeParent(*BACnetNotificationParametersChangeOfValueNewValue, *BACnetOpeningTag, *BACnetTagHeader, *BACnetClosingTag)
		GetParent() *BACnetNotificationParametersChangeOfValueNewValue
	}
	var _child BACnetNotificationParametersChangeOfValueNewValueChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfValueNewValueChangedBits
		_child, typeSwitchError = BACnetNotificationParametersChangeOfValueNewValueChangedBitsParse(readBuffer, tagNumber, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfValueNewValueChangedValue
		_child, typeSwitchError = BACnetNotificationParametersChangeOfValueNewValueChangedValueParse(readBuffer, tagNumber, peekedTagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValue"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), openingTag, peekedTagHeader, closingTag)
	return _child.GetParent(), nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParametersChangeOfValueNewValue, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValue"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}