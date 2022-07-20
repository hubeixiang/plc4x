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

// SALDataMetering is the corresponding interface of SALDataMetering
type SALDataMetering interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetMeteringData returns MeteringData (property field)
	GetMeteringData() MeteringData
}

// SALDataMeteringExactly can be used when we want exactly this type and not a type which fulfills SALDataMetering.
// This is useful for switch cases.
type SALDataMeteringExactly interface {
	SALDataMetering
	isSALDataMetering() bool
}

// _SALDataMetering is the data-structure of this message
type _SALDataMetering struct {
	*_SALData
	MeteringData MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataMetering) GetApplicationId() ApplicationId {
	return ApplicationId_METERING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataMetering) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataMetering) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataMetering) GetMeteringData() MeteringData {
	return m.MeteringData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataMetering factory function for _SALDataMetering
func NewSALDataMetering(meteringData MeteringData, salData SALData) *_SALDataMetering {
	_result := &_SALDataMetering{
		MeteringData: meteringData,
		_SALData:     NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataMetering(structType interface{}) SALDataMetering {
	if casted, ok := structType.(SALDataMetering); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataMetering); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataMetering) GetTypeName() string {
	return "SALDataMetering"
}

func (m *_SALDataMetering) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataMetering) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (meteringData)
	lengthInBits += m.MeteringData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataMetering) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataMeteringParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataMetering, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataMetering"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataMetering")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (meteringData)
	if pullErr := readBuffer.PullContext("meteringData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for meteringData")
	}
	_meteringData, _meteringDataErr := MeteringDataParse(readBuffer)
	if _meteringDataErr != nil {
		return nil, errors.Wrap(_meteringDataErr, "Error parsing 'meteringData' field of SALDataMetering")
	}
	meteringData := _meteringData.(MeteringData)
	if closeErr := readBuffer.CloseContext("meteringData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for meteringData")
	}

	if closeErr := readBuffer.CloseContext("SALDataMetering"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataMetering")
	}

	// Create a partially initialized instance
	_child := &_SALDataMetering{
		MeteringData: meteringData,
		_SALData:     &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataMetering) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataMetering"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataMetering")
		}

		// Simple Field (meteringData)
		if pushErr := writeBuffer.PushContext("meteringData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for meteringData")
		}
		_meteringDataErr := writeBuffer.WriteSerializable(m.GetMeteringData())
		if popErr := writeBuffer.PopContext("meteringData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for meteringData")
		}
		if _meteringDataErr != nil {
			return errors.Wrap(_meteringDataErr, "Error serializing 'meteringData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataMetering"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataMetering")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataMetering) isSALDataMetering() bool {
	return true
}

func (m *_SALDataMetering) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}