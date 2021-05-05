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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AmsPacket struct {
	TargetAmsNetId *AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId *AmsNetId
	SourceAmsPort  uint16
	CommandId      CommandId
	State          *State
	ErrorCode      uint32
	InvokeId       uint32
	Data           *AdsData
}

// The corresponding interface
type IAmsPacket interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewAmsPacket(targetAmsNetId *AmsNetId, targetAmsPort uint16, sourceAmsNetId *AmsNetId, sourceAmsPort uint16, commandId CommandId, state *State, errorCode uint32, invokeId uint32, data *AdsData) *AmsPacket {
	return &AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, CommandId: commandId, State: state, ErrorCode: errorCode, InvokeId: invokeId, Data: data}
}

func CastAmsPacket(structType interface{}) *AmsPacket {
	castFunc := func(typ interface{}) *AmsPacket {
		if casted, ok := typ.(AmsPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *AmsPacket) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AmsPacket) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.LengthInBits()

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.LengthInBits()

	// Simple field (sourceAmsPort)
	lengthInBits += 16

	// Simple field (commandId)
	lengthInBits += 16

	// Simple field (state)
	lengthInBits += m.State.LengthInBits()

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.LengthInBits()

	return lengthInBits
}

func (m *AmsPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsPacketParse(io utils.ReadBuffer) (*AmsPacket, error) {
	if pullErr := io.PullContext("AmsPacket"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := io.PullContext("targetAmsNetId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (targetAmsNetId)
	targetAmsNetId, _targetAmsNetIdErr := AmsNetIdParse(io)
	if _targetAmsNetIdErr != nil {
		return nil, errors.Wrap(_targetAmsNetIdErr, "Error parsing 'targetAmsNetId' field")
	}
	if closeErr := io.CloseContext("targetAmsNetId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (targetAmsPort)
	targetAmsPort, _targetAmsPortErr := io.ReadUint16("targetAmsPort", 16)
	if _targetAmsPortErr != nil {
		return nil, errors.Wrap(_targetAmsPortErr, "Error parsing 'targetAmsPort' field")
	}

	if pullErr := io.PullContext("sourceAmsNetId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (sourceAmsNetId)
	sourceAmsNetId, _sourceAmsNetIdErr := AmsNetIdParse(io)
	if _sourceAmsNetIdErr != nil {
		return nil, errors.Wrap(_sourceAmsNetIdErr, "Error parsing 'sourceAmsNetId' field")
	}
	if closeErr := io.CloseContext("sourceAmsNetId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort, _sourceAmsPortErr := io.ReadUint16("sourceAmsPort", 16)
	if _sourceAmsPortErr != nil {
		return nil, errors.Wrap(_sourceAmsPortErr, "Error parsing 'sourceAmsPort' field")
	}

	if pullErr := io.PullContext("commandId"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (commandId)
	commandId, _commandIdErr := CommandIdParse(io)
	if _commandIdErr != nil {
		return nil, errors.Wrap(_commandIdErr, "Error parsing 'commandId' field")
	}
	if closeErr := io.CloseContext("commandId"); closeErr != nil {
		return nil, closeErr
	}

	if pullErr := io.PullContext("state"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (state)
	state, _stateErr := StateParse(io)
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field")
	}
	if closeErr := io.CloseContext("state"); closeErr != nil {
		return nil, closeErr
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := io.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode, _errorCodeErr := io.ReadUint32("errorCode", 32)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId, _invokeIdErr := io.ReadUint32("invokeId", 32)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field")
	}

	if pullErr := io.PullContext("data"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (data)
	data, _dataErr := AdsDataParse(io, &commandId, state.Response)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	if closeErr := io.CloseContext("data"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("AmsPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, commandId, state, errorCode, invokeId, data), nil
}

func (m *AmsPacket) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("AmsPacket"); pushErr != nil {
		return pushErr
	}

	// Simple Field (targetAmsNetId)
	if pushErr := io.PushContext("targetAmsNetId"); pushErr != nil {
		return pushErr
	}
	_targetAmsNetIdErr := m.TargetAmsNetId.Serialize(io)
	if popErr := io.PopContext("targetAmsNetId"); popErr != nil {
		return popErr
	}
	if _targetAmsNetIdErr != nil {
		return errors.Wrap(_targetAmsNetIdErr, "Error serializing 'targetAmsNetId' field")
	}

	// Simple Field (targetAmsPort)
	targetAmsPort := uint16(m.TargetAmsPort)
	_targetAmsPortErr := io.WriteUint16("targetAmsPort", 16, (targetAmsPort))
	if _targetAmsPortErr != nil {
		return errors.Wrap(_targetAmsPortErr, "Error serializing 'targetAmsPort' field")
	}

	// Simple Field (sourceAmsNetId)
	if pushErr := io.PushContext("sourceAmsNetId"); pushErr != nil {
		return pushErr
	}
	_sourceAmsNetIdErr := m.SourceAmsNetId.Serialize(io)
	if popErr := io.PopContext("sourceAmsNetId"); popErr != nil {
		return popErr
	}
	if _sourceAmsNetIdErr != nil {
		return errors.Wrap(_sourceAmsNetIdErr, "Error serializing 'sourceAmsNetId' field")
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort := uint16(m.SourceAmsPort)
	_sourceAmsPortErr := io.WriteUint16("sourceAmsPort", 16, (sourceAmsPort))
	if _sourceAmsPortErr != nil {
		return errors.Wrap(_sourceAmsPortErr, "Error serializing 'sourceAmsPort' field")
	}

	// Simple Field (commandId)
	if pushErr := io.PushContext("commandId"); pushErr != nil {
		return pushErr
	}
	_commandIdErr := m.CommandId.Serialize(io)
	if popErr := io.PopContext("commandId"); popErr != nil {
		return popErr
	}
	if _commandIdErr != nil {
		return errors.Wrap(_commandIdErr, "Error serializing 'commandId' field")
	}

	// Simple Field (state)
	if pushErr := io.PushContext("state"); pushErr != nil {
		return pushErr
	}
	_stateErr := m.State.Serialize(io)
	if popErr := io.PopContext("state"); popErr != nil {
		return popErr
	}
	if _stateErr != nil {
		return errors.Wrap(_stateErr, "Error serializing 'state' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.Data.LengthInBytes())
	_lengthErr := io.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode := uint32(m.ErrorCode)
	_errorCodeErr := io.WriteUint32("errorCode", 32, (errorCode))
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId := uint32(m.InvokeId)
	_invokeIdErr := io.WriteUint32("invokeId", 32, (invokeId))
	if _invokeIdErr != nil {
		return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
	}

	// Simple Field (data)
	if pushErr := io.PushContext("data"); pushErr != nil {
		return pushErr
	}
	_dataErr := m.Data.Serialize(io)
	if popErr := io.PopContext("data"); popErr != nil {
		return popErr
	}
	if _dataErr != nil {
		return errors.Wrap(_dataErr, "Error serializing 'data' field")
	}

	if popErr := io.PopContext("AmsPacket"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *AmsPacket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
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
			case "targetAmsNetId":
				var data AmsNetId
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TargetAmsNetId = &data
			case "targetAmsPort":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TargetAmsPort = data
			case "sourceAmsNetId":
				var data AmsNetId
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAmsNetId = &data
			case "sourceAmsPort":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAmsPort = data
			case "commandId":
				var data CommandId
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommandId = data
			case "state":
				var data State
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.State = &data
			case "errorCode":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ErrorCode = data
			case "invokeId":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.InvokeId = data
			case "data":
				var dt *AdsData
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Data = dt
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *AmsPacket) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.AmsPacket"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TargetAmsNetId, xml.StartElement{Name: xml.Name{Local: "targetAmsNetId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TargetAmsPort, xml.StartElement{Name: xml.Name{Local: "targetAmsPort"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAmsNetId, xml.StartElement{Name: xml.Name{Local: "sourceAmsNetId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAmsPort, xml.StartElement{Name: xml.Name{Local: "sourceAmsPort"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.CommandId, xml.StartElement{Name: xml.Name{Local: "commandId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.State, xml.StartElement{Name: xml.Name{Local: "state"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ErrorCode, xml.StartElement{Name: xml.Name{Local: "errorCode"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.InvokeId, xml.StartElement{Name: xml.Name{Local: "invokeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m AmsPacket) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m AmsPacket) Box(name string, width int) utils.AsciiBox {
	boxName := "AmsPacket"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Complex field (case complex)
	boxes = append(boxes, m.TargetAmsNetId.Box("targetAmsNetId", width-2))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TargetAmsPort", m.TargetAmsPort, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.SourceAmsNetId.Box("sourceAmsNetId", width-2))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("SourceAmsPort", m.SourceAmsPort, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.CommandId.Box("commandId", width-2))
	// Complex field (case complex)
	boxes = append(boxes, m.State.Box("state", width-2))
	// Implicit Field (length)
	length := uint32(m.Data.LengthInBytes())
	// uint32 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Length", length, -1))
	// Simple field (case simple)
	// uint32 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ErrorCode", m.ErrorCode, -1))
	// Simple field (case simple)
	// uint32 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("InvokeId", m.InvokeId, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Data.Box("data", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}