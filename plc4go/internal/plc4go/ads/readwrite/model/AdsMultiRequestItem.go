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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsMultiRequestItem struct {
	Child IAdsMultiRequestItemChild
}

// The corresponding interface
type IAdsMultiRequestItem interface {
	IndexGroup() uint32
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IAdsMultiRequestItemParent interface {
	SerializeParent(io utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsMultiRequestItemChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *AdsMultiRequestItem)
	GetTypeName() string
	IAdsMultiRequestItem
	utils.AsciiBoxer
}

func NewAdsMultiRequestItem() *AdsMultiRequestItem {
	return &AdsMultiRequestItem{}
}

func CastAdsMultiRequestItem(structType interface{}) *AdsMultiRequestItem {
	castFunc := func(typ interface{}) *AdsMultiRequestItem {
		if casted, ok := typ.(AdsMultiRequestItem); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsMultiRequestItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *AdsMultiRequestItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsMultiRequestItem) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *AdsMultiRequestItem) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsMultiRequestItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsMultiRequestItemParse(io utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItem, error) {
	if pullErr := io.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *AdsMultiRequestItem
	var typeSwitchError error
	switch {
	case indexGroup == 61568: // AdsMultiRequestItemRead
		_parent, typeSwitchError = AdsMultiRequestItemReadParse(io)
	case indexGroup == 61569: // AdsMultiRequestItemWrite
		_parent, typeSwitchError = AdsMultiRequestItemWriteParse(io)
	case indexGroup == 61570: // AdsMultiRequestItemReadWrite
		_parent, typeSwitchError = AdsMultiRequestItemReadWriteParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *AdsMultiRequestItem) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *AdsMultiRequestItem) SerializeParent(io utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("AdsMultiRequestItem"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *AdsMultiRequestItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of AdsMultiRequestItem")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemRead":
					var dt *AdsMultiRequestItemRead
					if m.Child != nil {
						dt = m.Child.(*AdsMultiRequestItemRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemWrite":
					var dt *AdsMultiRequestItemWrite
					if m.Child != nil {
						dt = m.Child.(*AdsMultiRequestItemWrite)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.ads.readwrite.AdsMultiRequestItemReadWrite":
					var dt *AdsMultiRequestItemReadWrite
					if m.Child != nil {
						dt = m.Child.(*AdsMultiRequestItemReadWrite)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *AdsMultiRequestItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.ads.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m AdsMultiRequestItem) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *AdsMultiRequestItem) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *AdsMultiRequestItem) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "AdsMultiRequestItem"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}