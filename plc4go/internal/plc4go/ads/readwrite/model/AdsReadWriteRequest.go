//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsReadWriteRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	ReadLength  uint32
	Items       []*AdsMultiRequestItem
	Data        []int8
	Parent      *AdsData
}

// The corresponding interface
type IAdsReadWriteRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadWriteRequest) CommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *AdsReadWriteRequest) Response() bool {
	return false
}

func (m *AdsReadWriteRequest) InitializeParent(parent *AdsData) {
}

func NewAdsReadWriteRequest(indexGroup uint32, indexOffset uint32, readLength uint32, items []*AdsMultiRequestItem, data []int8) *AdsData {
	child := &AdsReadWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadWriteRequest(structType interface{}) *AdsReadWriteRequest {
	castFunc := func(typ interface{}) *AdsReadWriteRequest {
		if casted, ok := typ.(AdsReadWriteRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadWriteRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadWriteRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadWriteRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadWriteRequest) GetTypeName() string {
	return "AdsReadWriteRequest"
}

func (m *AdsReadWriteRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsReadWriteRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (readLength)
	lengthInBits += 32

	// Implicit Field (writeLength)
	lengthInBits += 32

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsReadWriteRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadWriteRequestParse(io utils.ReadBuffer) (*AdsData, error) {
	if pullErr := io.PullContext("AdsReadWriteRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (indexGroup)
	indexGroup, _indexGroupErr := io.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}

	// Simple Field (indexOffset)
	indexOffset, _indexOffsetErr := io.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}

	// Simple Field (readLength)
	readLength, _readLengthErr := io.ReadUint32("readLength", 32)
	if _readLengthErr != nil {
		return nil, errors.Wrap(_readLengthErr, "Error parsing 'readLength' field")
	}

	// Implicit Field (writeLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	writeLength, _writeLengthErr := io.ReadUint32("writeLength", 32)
	_ = writeLength
	if _writeLengthErr != nil {
		return nil, errors.Wrap(_writeLengthErr, "Error parsing 'writeLength' field")
	}

	// Array field (items)
	if pullErr := io.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*AdsMultiRequestItem, utils.InlineIf(bool(bool(bool(bool(bool((indexGroup) == (61568)))) || bool(bool(bool((indexGroup) == (61569))))) || bool(bool(bool((indexGroup) == (61570))))), func() uint16 { return uint16(indexOffset) }, func() uint16 { return uint16(uint16(0)) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool(bool(bool(bool((indexGroup) == (61568)))) || bool(bool(bool((indexGroup) == (61569))))) || bool(bool(bool((indexGroup) == (61570))))), func() uint16 { return uint16(indexOffset) }, func() uint16 { return uint16(uint16(0)) })); curItem++ {
		_item, _err := AdsMultiRequestItemParse(io, indexGroup)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	if closeErr := io.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Array field (data)
	if pullErr := io.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]int8, uint16(writeLength)-uint16(uint16(uint16(uint16(len(items)))*uint16(uint16(12)))))
	for curItem := uint16(0); curItem < uint16(uint16(writeLength)-uint16(uint16(uint16(uint16(len(items)))*uint16(uint16(12))))); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := io.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("AdsReadWriteRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadWriteRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("AdsReadWriteRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := io.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := io.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (readLength)
		readLength := uint32(m.ReadLength)
		_readLengthErr := io.WriteUint32("readLength", 32, (readLength))
		if _readLengthErr != nil {
			return errors.Wrap(_readLengthErr, "Error serializing 'readLength' field")
		}

		// Implicit Field (writeLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		writeLength := uint32(uint32(uint32(uint32(uint32(len(m.Items)))*uint32(uint32(utils.InlineIf(bool(bool((m.IndexGroup) == (61570))), func() uint16 { return uint16(uint32(16)) }, func() uint16 { return uint16(uint32(12)) }))))) + uint32(uint32(len(m.Data))))
		_writeLengthErr := io.WriteUint32("writeLength", 32, (writeLength))
		if _writeLengthErr != nil {
			return errors.Wrap(_writeLengthErr, "Error serializing 'writeLength' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := io.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := io.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := io.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := io.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("AdsReadWriteRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *AdsReadWriteRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "indexGroup":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexGroup = data
			case "indexOffset":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IndexOffset = data
			case "readLength":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReadLength = data
			case "items":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *AdsMultiRequestItem
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.Items = append(m.Items, dt)
					default:
						break arrayLoop
					}
				}
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *AdsReadWriteRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.IndexGroup, xml.StartElement{Name: xml.Name{Local: "indexGroup"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IndexOffset, xml.StartElement{Name: xml.Name{Local: "indexOffset"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReadLength, xml.StartElement{Name: xml.Name{Local: "readLength"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Items {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsReadWriteRequest) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m AdsReadWriteRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsReadWriteRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("IndexGroup", m.IndexGroup, -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("IndexOffset", m.IndexOffset, -1))
		// Simple field (case simple)
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ReadLength", m.ReadLength, -1))
		// Implicit Field (writeLength)
		writeLength := uint32(uint32(uint32(uint32(uint32(len(m.Items)))*uint32(uint32(utils.InlineIf(bool(bool((m.IndexGroup) == (61570))), func() uint16 { return uint16(uint32(16)) }, func() uint16 { return uint16(uint32(12)) }))))) + uint32(uint32(len(m.Data))))
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("WriteLength", writeLength, -1))
		// Array Field (items)
		if m.Items != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Items {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Items", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}