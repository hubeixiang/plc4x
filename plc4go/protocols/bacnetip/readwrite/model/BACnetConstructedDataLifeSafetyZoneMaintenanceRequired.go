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

// BACnetConstructedDataLifeSafetyZoneMaintenanceRequired is the corresponding interface of BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
type BACnetConstructedDataLifeSafetyZoneMaintenanceRequired interface {
	BACnetConstructedData
	// GetMaintenanceRequired returns MaintenanceRequired (property field)
	GetMaintenanceRequired() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired struct {
	*_BACnetConstructedData
	MaintenanceRequired BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAINTENANCE_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetMaintenanceRequired() BACnetApplicationTagBoolean {
	return m.MaintenanceRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetMaintenanceRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequired factory function for _BACnetConstructedDataLifeSafetyZoneMaintenanceRequired
func NewBACnetConstructedDataLifeSafetyZoneMaintenanceRequired(maintenanceRequired BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	_result := &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired{
		MaintenanceRequired:    maintenanceRequired,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZoneMaintenanceRequired(structType interface{}) BACnetConstructedDataLifeSafetyZoneMaintenanceRequired {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZoneMaintenanceRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZoneMaintenanceRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maintenanceRequired)
	lengthInBits += m.MaintenanceRequired.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyZoneMaintenanceRequiredParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLifeSafetyZoneMaintenanceRequired, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maintenanceRequired)
	if pullErr := readBuffer.PullContext("maintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maintenanceRequired")
	}
	_maintenanceRequired, _maintenanceRequiredErr := BACnetApplicationTagParse(readBuffer)
	if _maintenanceRequiredErr != nil {
		return nil, errors.Wrap(_maintenanceRequiredErr, "Error parsing 'maintenanceRequired' field")
	}
	maintenanceRequired := _maintenanceRequired.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("maintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maintenanceRequired")
	}

	// Virtual field
	_actualValue := maintenanceRequired
	actualValue := _actualValue.(BACnetApplicationTagBoolean)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired{
		MaintenanceRequired:    maintenanceRequired,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
		}

		// Simple Field (maintenanceRequired)
		if pushErr := writeBuffer.PushContext("maintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maintenanceRequired")
		}
		_maintenanceRequiredErr := writeBuffer.WriteSerializable(m.GetMaintenanceRequired())
		if popErr := writeBuffer.PopContext("maintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maintenanceRequired")
		}
		if _maintenanceRequiredErr != nil {
			return errors.Wrap(_maintenanceRequiredErr, "Error serializing 'maintenanceRequired' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZoneMaintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZoneMaintenanceRequired")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZoneMaintenanceRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}