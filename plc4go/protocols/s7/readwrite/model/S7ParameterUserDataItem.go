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

// S7ParameterUserDataItem is the corresponding interface of S7ParameterUserDataItem
type S7ParameterUserDataItem interface {
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _S7ParameterUserDataItem is the data-structure of this message
type _S7ParameterUserDataItem struct {
	_S7ParameterUserDataItemChildRequirements
}

type _S7ParameterUserDataItemChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetItemType() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

type S7ParameterUserDataItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child S7ParameterUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7ParameterUserDataItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent S7ParameterUserDataItem)
	GetParent() *S7ParameterUserDataItem

	GetTypeName() string
	S7ParameterUserDataItem
}

// NewS7ParameterUserDataItem factory function for _S7ParameterUserDataItem
func NewS7ParameterUserDataItem() *_S7ParameterUserDataItem {
	return &_S7ParameterUserDataItem{}
}

// Deprecated: use the interface for direct cast
func CastS7ParameterUserDataItem(structType interface{}) S7ParameterUserDataItem {
	if casted, ok := structType.(S7ParameterUserDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterUserDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterUserDataItem) GetTypeName() string {
	return "S7ParameterUserDataItem"
}

func (m *_S7ParameterUserDataItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterUserDataItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterUserDataItemParse(readBuffer utils.ReadBuffer) (S7ParameterUserDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterUserDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterUserDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint8("itemType", 8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7ParameterUserDataItemChildSerializeRequirement interface {
		S7ParameterUserDataItem
		InitializeParent(S7ParameterUserDataItem)
		GetParent() S7ParameterUserDataItem
	}
	var _childTemp interface{}
	var _child S7ParameterUserDataItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7ParameterUserDataItemCPUFunctions
		_childTemp, typeSwitchError = S7ParameterUserDataItemCPUFunctionsParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(S7ParameterUserDataItemChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7ParameterUserDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterUserDataItem")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_S7ParameterUserDataItem) SerializeParent(writeBuffer utils.WriteBuffer, child S7ParameterUserDataItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7ParameterUserDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7ParameterUserDataItem")
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.GetItemType())
	_itemTypeErr := writeBuffer.WriteUint8("itemType", 8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7ParameterUserDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7ParameterUserDataItem")
	}
	return nil
}

func (m *_S7ParameterUserDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
