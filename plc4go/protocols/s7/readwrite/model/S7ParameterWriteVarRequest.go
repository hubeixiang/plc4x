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

// S7ParameterWriteVarRequest is the data-structure of this message
type S7ParameterWriteVarRequest struct {
	*S7Parameter
	Items []*S7VarRequestParameterItem
}

// IS7ParameterWriteVarRequest is the corresponding interface of S7ParameterWriteVarRequest
type IS7ParameterWriteVarRequest interface {
	IS7Parameter
	// GetItems returns Items (property field)
	GetItems() []*S7VarRequestParameterItem
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

func (m *S7ParameterWriteVarRequest) GetParameterType() uint8 {
	return 0x05
}

func (m *S7ParameterWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7ParameterWriteVarRequest) InitializeParent(parent *S7Parameter) {}

func (m *S7ParameterWriteVarRequest) GetParent() *S7Parameter {
	return m.S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7ParameterWriteVarRequest) GetItems() []*S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterWriteVarRequest factory function for S7ParameterWriteVarRequest
func NewS7ParameterWriteVarRequest(items []*S7VarRequestParameterItem) *S7ParameterWriteVarRequest {
	_result := &S7ParameterWriteVarRequest{
		Items:       items,
		S7Parameter: NewS7Parameter(),
	}
	_result.Child = _result
	return _result
}

func CastS7ParameterWriteVarRequest(structType interface{}) *S7ParameterWriteVarRequest {
	if casted, ok := structType.(S7ParameterWriteVarRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(S7Parameter); ok {
		return CastS7ParameterWriteVarRequest(casted.Child)
	}
	if casted, ok := structType.(*S7Parameter); ok {
		return CastS7ParameterWriteVarRequest(casted.Child)
	}
	return nil
}

func (m *S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *S7ParameterWriteVarRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7ParameterWriteVarRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7ParameterWriteVarRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterWriteVarRequestParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7ParameterWriteVarRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (numItems) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numItems, _numItemsErr := readBuffer.ReadUint8("numItems", 8)
	_ = numItems
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7VarRequestParameterItem, numItems)
	{
		for curItem := uint16(0); curItem < uint16(numItems); curItem++ {
			_item, _err := S7VarRequestParameterItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastS7VarRequestParameterItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7ParameterWriteVarRequest{
		Items:       items,
		S7Parameter: &S7Parameter{},
	}
	_child.S7Parameter.Child = _child
	return _child, nil
}

func (m *S7ParameterWriteVarRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarRequest"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (numItems) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numItems := uint8(uint8(len(m.GetItems())))
		_numItemsErr := writeBuffer.WriteUint8("numItems", 8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7ParameterWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}