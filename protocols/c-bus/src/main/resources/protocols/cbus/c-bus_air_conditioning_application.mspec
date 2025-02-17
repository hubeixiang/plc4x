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

[type AirConditioningData
    //TODO: golang doesn't like checking for null so we use that static call to check that the enum is known
    [validation 'STATIC_CALL("knowsAirConditioningCommandTypeContainer", readBuffer)' "no command type could be found" shouldFail=false]
    [simple  AirConditioningCommandTypeContainer      commandTypeContainer                                   ]
    [virtual AirConditioningCommandType               commandType          'commandTypeContainer.commandType']
    [typeSwitch commandType
        ['HVAC_SCHEDULE_ENTRY'              *HvacScheduleEntry
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   uint 8              entry               ]
            [simple   byte                format              ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
            [simple   HVACStartTime       startTime           ]
            [optional HVACTemperature     level 'hvacModeAndFlags.isLevelTemperature'   ]
            [optional HVACRawLevels       rawLevel 'hvacModeAndFlags.isLevelRaw'        ]
        ]
        ['HUMIDITY_SCHEDULE_ENTRY'          *HumidityScheduleEntry
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   uint 8                    entry                   ]
            [simple   byte                      format                  ]
            [simple   HVACHumidityModeAndFlags  humidityModeAndFlags    ]
            [simple   HVACStartTime             startTime               ]
            [optional HVACHumidity              level    'humidityModeAndFlags.isLevelHumidity'   ]
            [optional HVACRawLevels             rawLevel 'humidityModeAndFlags.isLevelRaw'        ]
        ]
        ['REFRESH'                          *Refresh
            [simple   byte                      zoneGroup               ]
        ]
        ['ZONE_HVAC_PLANT_STATUS'           *ZoneHvacPlantStatus
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACType                  hvacType                ]
            [simple   HVACStatusFlags           hvacStatus              ]
            [simple   HVACError                 hvacErrorCode           ]
        ]
        ['ZONE_HUMIDITY_PLANT_STATUS'       *ZoneHumidityPlantStatus
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACHumidityType          humidityType            ]
            [simple   HVACHumidityStatusFlags   humidityStatus          ]
            [simple   HVACHumidityError         humidityErrorCode       ]
        ]
        ['ZONE_TEMPERATURE'                 *ZoneTemperature
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACTemperature           temperature             ]
            [simple   HVACSensorStatus          sensorStatus            ]
        ]
        ['ZONE_HUMIDITY'                    *ZoneHumidity
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACHumidity              humidity                ]
            [simple   HVACSensorStatus          sensorStatus            ]
        ]
        ['SET_ZONE_GROUP_OFF'               *SetZoneGroupOff
            [simple   byte                      zoneGroup               ]
        ]
        ['SET_ZONE_GROUP_ON'                *SetZoneGroupOn
            [simple   byte                      zoneGroup               ]
        ]
        ['SET_ZONE_HVAC_MODE'               *SetZoneHvacMode
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
            [simple   HVACType            hvacType            ]
            [optional HVACTemperature     level     'hvacModeAndFlags.isLevelTemperature']
            [optional HVACRawLevels       rawLevel  'hvacModeAndFlags.isLevelRaw'        ]
            [optional HVACAuxiliaryLevel  auxLevel  'hvacModeAndFlags.isAuxLevelUsed'    ]
        ]
        ['SET_PLANT_HVAC_LEVEL'             *SetPlantHvacLevel
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
            [simple   HVACType            hvacType            ]
            [optional HVACTemperature     level     'hvacModeAndFlags.isLevelTemperature']
            [optional HVACRawLevels       rawLevel  'hvacModeAndFlags.isLevelRaw'        ]
            [optional HVACAuxiliaryLevel  auxLevel  'hvacModeAndFlags.isAuxLevelUsed'    ]
        ]
        ['SET_ZONE_HUMIDITY_MODE'           *SetZoneHumidityMode
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACHumidityModeAndFlags  humidityModeAndFlags    ]
            [simple   HVACHumidityType          humidityType            ]
            [optional HVACHumidity              level    'humidityModeAndFlags.isLevelHumidity']
            [optional HVACRawLevels             rawLevel 'humidityModeAndFlags.isLevelRaw'     ]
            [optional HVACAuxiliaryLevel        auxLevel 'humidityModeAndFlags.isAuxLevelUsed' ]
        ]
        ['SET_PLANT_HUMIDITY_LEVEL'         *SetPlantHumidityLevel
            [simple   byte                      zoneGroup               ]
            [simple   HVACZoneList              zoneList                ]
            [simple   HVACHumidityModeAndFlags  humidityModeAndFlags    ]
            [simple   HVACHumidityType          humidityType            ]
            [optional HVACHumidity              level    'humidityModeAndFlags.isLevelHumidity']
            [optional HVACRawLevels             rawLevel 'humidityModeAndFlags.isLevelRaw'     ]
            [optional HVACAuxiliaryLevel        auxLevel 'humidityModeAndFlags.isAuxLevelUsed' ]
        ]
        ['SET_HVAC_UPPER_GUARD_LIMIT'       *SetHvacUpperGuardLimit
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   HVACTemperature     limit               ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
        ]
        ['SET_HVAC_LOWER_GUARD_LIMIT'       *SetHvacLowerGuardLimit
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   HVACTemperature     limit               ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
        ]
        ['SET_HVAC_SETBACK_LIMIT'           *SetHvacSetbackLimit
            [simple   byte                zoneGroup           ]
            [simple   HVACZoneList        zoneList            ]
            [simple   HVACTemperature     limit               ]
            [simple   HVACModeAndFlags    hvacModeAndFlags    ]
        ]
        ['SET_HUMIDITY_UPPER_GUARD_LIMIT'   *SetHumidityUpperGuardLimit
            [simple   byte                      zoneGroup           ]
            [simple   HVACZoneList              zoneList            ]
            [simple   HVACHumidity              limit               ]
            [simple   HVACHumidityModeAndFlags  hvacModeAndFlags    ]
        ]
        ['SET_HUMIDITY_LOWER_GUARD_LIMIT'   *SetHumidityLowerGuardLimit
            [simple   byte                      zoneGroup           ]
            [simple   HVACZoneList              zoneList            ]
            [simple   HVACHumidity              limit               ]
            [simple   HVACHumidityModeAndFlags  hvacModeAndFlags    ]
        ]
        ['SET_HUMIDITY_SETBACK_LIMIT'       *SetHumiditySetbackLimit
            [simple   byte                      zoneGroup           ]
            [simple   HVACZoneList              zoneList            ]
            [simple   HVACHumidity              limit               ]
            [simple   HVACHumidityModeAndFlags  hvacModeAndFlags    ]
        ]
    ]
]

[enum uint 8 AirConditioningCommandTypeContainer(AirConditioningCommandType commandType, uint 5 numBytes)
    ['0x01' AirConditioningCommandSetZoneGroupOff               ['SET_ZONE_GROUP_OFF',              '1']]
    ['0x05' AirConditioningCommandZoneHvacPlantStatus           ['ZONE_HVAC_PLANT_STATUS',          '5']]
    ['0x0D' AirConditioningCommandZoneHumidityPlantStatus       ['ZONE_HUMIDITY_PLANT_STATUS',      '5']]
    ['0x15' AirConditioningCommandZoneTemperature               ['ZONE_TEMPERATURE',                '5']]
    ['0x1D' AirConditioningCommandZoneHumidity                  ['ZONE_HUMIDITY',                   '5']]
    ['0x21' AirConditioningCommandRefresh                       ['REFRESH',                         '1']]
    ['0x2F' AirConditioningCommandSetZoneHvacMode               ['SET_ZONE_HVAC_MODE',              '7']]
    ['0x36' AirConditioningCommandSetPlantHvacLevel             ['SET_PLANT_HVAC_LEVEL',            '6']]
    ['0x47' AirConditioningCommandSetZoneHumidityMode           ['SET_ZONE_HUMIDITY_MODE',          '7']]
    ['0x4E' AirConditioningCommandSetPlantHumidityLevel         ['SET_PLANT_HUMIDITY_LEVEL',        '6']]
    ['0x55' AirConditioningCommandSetHvacUpperGuardLimit        ['SET_HVAC_UPPER_GUARD_LIMIT',      '5']]
    ['0x5D' AirConditioningCommandSetHvacLowerGuardLimit        ['SET_HVAC_LOWER_GUARD_LIMIT',      '5']]
    ['0x65' AirConditioningCommandSetHvacSetbackLimit           ['SET_HVAC_SETBACK_LIMIT',          '5']]
    ['0x6D' AirConditioningCommandSetHumidityUpperGuardLimit    ['SET_HUMIDITY_UPPER_GUARD_LIMIT',  '5']]
    ['0x75' AirConditioningCommandSetHumidityLowerGuardLimit    ['SET_HUMIDITY_LOWER_GUARD_LIMIT',  '5']]
    ['0x79' AirConditioningCommandSetZoneGroupOn                ['SET_ZONE_GROUP_ON',               '1']]
    ['0x7D' AirConditioningCommandSetHumiditySetbackLimit       ['SET_HUMIDITY_SETBACK_LIMIT',      '5']]
    ['0x89' AirConditioningCommandHvacScheduleEntry             ['HVAC_SCHEDULE_ENTRY',             '9']]
    ['0xA9' AirConditioningCommandHumidityScheduleEntry         ['HUMIDITY_SCHEDULE_ENTRY',         '9']]
]

[enum uint 4 AirConditioningCommandType
    ['0x00' SET_ZONE_GROUP_OFF              ]
    ['0x01' ZONE_HVAC_PLANT_STATUS          ]
    ['0x02' ZONE_HUMIDITY_PLANT_STATUS      ]
    ['0x03' ZONE_TEMPERATURE                ]
    ['0x04' ZONE_HUMIDITY                   ]
    ['0x05' REFRESH                         ]
    ['0x06' SET_ZONE_HVAC_MODE              ]
    ['0x07' SET_PLANT_HVAC_LEVEL            ]
    ['0x08' SET_ZONE_HUMIDITY_MODE          ]
    ['0x09' SET_PLANT_HUMIDITY_LEVEL        ]
    ['0x0A' SET_HVAC_UPPER_GUARD_LIMIT      ]
    ['0x0B' SET_HVAC_LOWER_GUARD_LIMIT      ]
    ['0x0C' SET_HVAC_SETBACK_LIMIT          ]
    ['0x0D' SET_HUMIDITY_UPPER_GUARD_LIMIT  ]
    ['0x0E' SET_HUMIDITY_LOWER_GUARD_LIMIT  ]
    ['0x0F' SET_ZONE_GROUP_ON               ]
    ['0x10' SET_HUMIDITY_SETBACK_LIMIT      ]
    ['0x11' HVAC_SCHEDULE_ENTRY             ]
    ['0x12' HUMIDITY_SCHEDULE_ENTRY         ]
]

[type HVACTemperature
    // TODO: check values from Air Conditioning Application 25.5.1
    [simple   int 16    temperatureValue                            ]
    [virtual  float 32  temperatureInCelcius 'temperatureValue/256' ]
]

[type HVACHumidity
    // TODO: check values from Air Conditioning Application 25.5.2
    [simple   uint 16   humidityValue                               ]
    [virtual  float 32  humidityInPercent 'humidityValue/65535'     ]
]

[type HVACRawLevels
    // TODO: check values from Air Conditioning Application 25.5.3
    [simple   int 16    rawValue                                    ]
    [virtual  float 32  valueInPercent 'rawValue/32767'             ]
]

[type HVACModeAndFlags
    [reserved bit   'false'                                     ]
    [simple   bit   auxiliaryLevel                              ]
    [virtual  bit   isAuxLevelUnused '!auxiliaryLevel'          ]
    [virtual  bit   isAuxLevelUsed   'auxiliaryLevel'           ]
    [simple   bit   guard                                       ]
    [virtual  bit   isGuardDisabled '!guard'                    ]
    [virtual  bit   isGuardEnabled   'guard'                    ]
    [simple   bit   setback                                     ]
    [virtual  bit   isSetbackDisabled '!setback'                ]
    [virtual  bit   isSetbackEnabled   'setback'                ]
    [simple   bit   level                                       ]
    [virtual  bit   isLevelTemperature '!level'                 ]
    [virtual  bit   isLevelRaw         'level'                  ]
    [simple   HVACModeAndFlagsMode  mode                        ]
]

[enum uint 3 HVACModeAndFlagsMode
    ['0x0' OFF              ]
    ['0x1' HEAT_ONLY        ]
    ['0x2' COOL_ONLY        ]
    ['0x3' HEAT_AND_COOL    ]
    ['0x4' VENT_FAN_ONLY    ]
]

[enum uint 8 HVACType
    ['0x00' NONE                            ]
    ['0x01' FURNACE_GAS_OIL_ELECTRIC        ]
    ['0x02' EVAPORATIVE                     ]
    ['0x03' HEAT_PUMP_REVERSE_CYCLE         ]
    ['0x04' HEAT_PUMP_HEATING_ONLY          ]
    ['0x05' HEAT_PUMP_COOLING_ONLY          ]
    ['0x06' FURNANCE_EVAP_COOLING           ]
    ['0x07' FURNANCE_HEAT_PUMP_COOLING_ONLY ]
    ['0x08' HYDRONIC                        ]
    ['0x09' HYDRONIC_HEAT_PUMP_COOLING_ONLY ]
    ['0x0A' HYDRONIC_EVAPORATIVE            ]
    ['0xFF' ANY                             ]
]

[enum uint 8 HVACError
    ['0x00' NO_ERROR                    ]
    ['0x01' HEATER_TOTAL_FAILURE        ]
    ['0x02' COOLER_TOTAL_FAILURE        ]
    ['0x03' FAN_TOTAL_FAILURE           ]
    ['0x04' TEMPERATURE_SENSOR_FAILURE  ]
    ['0x05' HEATER_TEMPORARY_PROBLEM    ]
    ['0x06' COOLER_TEMPORARY_PROBLEM    ]
    ['0x07' FAN_TEMPORARY_PROBLEM       ]
    ['0x08' HEATER_SERVICE_REQUIRED     ]
    ['0x09' COOLER_SERVICE_REQUIRED     ]
    ['0x0A' FAN_SERVICE_REQUIRED        ]
    ['0x0B' FILTER_REPLACEMENT_REQUIRED ]
    ['0x80' CUSTOM_ERROR_0              ]
    ['0x81' CUSTOM_ERROR_1              ]
    ['0x82' CUSTOM_ERROR_2              ]
    ['0x83' CUSTOM_ERROR_3              ]
    ['0x84' CUSTOM_ERROR_4              ]
    ['0x85' CUSTOM_ERROR_5              ]
    ['0x86' CUSTOM_ERROR_6              ]
    ['0x87' CUSTOM_ERROR_7              ]
    ['0x88' CUSTOM_ERROR_8              ]
    ['0x89' CUSTOM_ERROR_9              ]
    ['0x8A' CUSTOM_ERROR_10             ]
    ['0x8B' CUSTOM_ERROR_11             ]
    ['0x8C' CUSTOM_ERROR_12             ]
    ['0x8D' CUSTOM_ERROR_13             ]
    ['0x8E' CUSTOM_ERROR_14             ]
    ['0x8F' CUSTOM_ERROR_15             ]
    ['0x90' CUSTOM_ERROR_16             ]
    ['0x91' CUSTOM_ERROR_17             ]
    ['0x92' CUSTOM_ERROR_18             ]
    ['0x93' CUSTOM_ERROR_19             ]
    ['0x94' CUSTOM_ERROR_20             ]
    ['0x95' CUSTOM_ERROR_21             ]
    ['0x96' CUSTOM_ERROR_22             ]
    ['0x97' CUSTOM_ERROR_23             ]
    ['0x98' CUSTOM_ERROR_24             ]
    ['0x99' CUSTOM_ERROR_25             ]
    ['0x9A' CUSTOM_ERROR_26             ]
    ['0x9B' CUSTOM_ERROR_27             ]
    ['0x9C' CUSTOM_ERROR_28             ]
    ['0x9D' CUSTOM_ERROR_29             ]
    ['0x9E' CUSTOM_ERROR_30             ]
    ['0x9F' CUSTOM_ERROR_31             ]
    ['0xA0' CUSTOM_ERROR_32             ]
    ['0xA1' CUSTOM_ERROR_33             ]
    ['0xA2' CUSTOM_ERROR_34             ]
    ['0xA3' CUSTOM_ERROR_35             ]
    ['0xA4' CUSTOM_ERROR_36             ]
    ['0xA5' CUSTOM_ERROR_37             ]
    ['0xA6' CUSTOM_ERROR_38             ]
    ['0xA7' CUSTOM_ERROR_39             ]
    ['0xA8' CUSTOM_ERROR_40             ]
    ['0xA9' CUSTOM_ERROR_41             ]
    ['0xAA' CUSTOM_ERROR_42             ]
    ['0xAB' CUSTOM_ERROR_43             ]
    ['0xAC' CUSTOM_ERROR_44             ]
    ['0xAD' CUSTOM_ERROR_45             ]
    ['0xAE' CUSTOM_ERROR_46             ]
    ['0xAF' CUSTOM_ERROR_47             ]
    ['0xB0' CUSTOM_ERROR_48             ]
    ['0xB1' CUSTOM_ERROR_49             ]
    ['0xB2' CUSTOM_ERROR_50             ]
    ['0xB3' CUSTOM_ERROR_51             ]
    ['0xB4' CUSTOM_ERROR_52             ]
    ['0xB5' CUSTOM_ERROR_53             ]
    ['0xB6' CUSTOM_ERROR_54             ]
    ['0xB7' CUSTOM_ERROR_55             ]
    ['0xB8' CUSTOM_ERROR_56             ]
    ['0xB9' CUSTOM_ERROR_57             ]
    ['0xBA' CUSTOM_ERROR_58             ]
    ['0xBB' CUSTOM_ERROR_59             ]
    ['0xBC' CUSTOM_ERROR_60             ]
    ['0xBD' CUSTOM_ERROR_61             ]
    ['0xBE' CUSTOM_ERROR_62             ]
    ['0xBF' CUSTOM_ERROR_63             ]
    ['0xC0' CUSTOM_ERROR_64             ]
    ['0xC1' CUSTOM_ERROR_65             ]
    ['0xC2' CUSTOM_ERROR_66             ]
    ['0xC3' CUSTOM_ERROR_67             ]
    ['0xC4' CUSTOM_ERROR_68             ]
    ['0xC5' CUSTOM_ERROR_69             ]
    ['0xC6' CUSTOM_ERROR_70             ]
    ['0xC7' CUSTOM_ERROR_71             ]
    ['0xC8' CUSTOM_ERROR_72             ]
    ['0xC9' CUSTOM_ERROR_73             ]
    ['0xCA' CUSTOM_ERROR_74             ]
    ['0xCB' CUSTOM_ERROR_75             ]
    ['0xCC' CUSTOM_ERROR_76             ]
    ['0xCD' CUSTOM_ERROR_77             ]
    ['0xCE' CUSTOM_ERROR_78             ]
    ['0xCF' CUSTOM_ERROR_79             ]
    ['0xD0' CUSTOM_ERROR_80             ]
    ['0xD1' CUSTOM_ERROR_81             ]
    ['0xD2' CUSTOM_ERROR_82             ]
    ['0xD3' CUSTOM_ERROR_83             ]
    ['0xD4' CUSTOM_ERROR_84             ]
    ['0xD5' CUSTOM_ERROR_85             ]
    ['0xD6' CUSTOM_ERROR_86             ]
    ['0xD7' CUSTOM_ERROR_87             ]
    ['0xD8' CUSTOM_ERROR_88             ]
    ['0xD9' CUSTOM_ERROR_89             ]
    ['0xDA' CUSTOM_ERROR_90             ]
    ['0xDB' CUSTOM_ERROR_91             ]
    ['0xDC' CUSTOM_ERROR_92             ]
    ['0xDD' CUSTOM_ERROR_93             ]
    ['0xDE' CUSTOM_ERROR_94             ]
    ['0xDF' CUSTOM_ERROR_95             ]
    ['0xE0' CUSTOM_ERROR_96             ]
    ['0xE1' CUSTOM_ERROR_97             ]
    ['0xE2' CUSTOM_ERROR_98             ]
    ['0xE3' CUSTOM_ERROR_99             ]
    ['0xE4' CUSTOM_ERROR_100            ]
    ['0xE5' CUSTOM_ERROR_101            ]
    ['0xE6' CUSTOM_ERROR_102            ]
    ['0xE7' CUSTOM_ERROR_103            ]
    ['0xE8' CUSTOM_ERROR_104            ]
    ['0xE9' CUSTOM_ERROR_105            ]
    ['0xEA' CUSTOM_ERROR_106            ]
    ['0xEB' CUSTOM_ERROR_107            ]
    ['0xEC' CUSTOM_ERROR_108            ]
    ['0xED' CUSTOM_ERROR_109            ]
    ['0xEE' CUSTOM_ERROR_110            ]
    ['0xEF' CUSTOM_ERROR_111            ]
    ['0xF0' CUSTOM_ERROR_112            ]
    ['0xF1' CUSTOM_ERROR_113            ]
    ['0xF2' CUSTOM_ERROR_114            ]
    ['0xF3' CUSTOM_ERROR_115            ]
    ['0xF4' CUSTOM_ERROR_116            ]
    ['0xF5' CUSTOM_ERROR_117            ]
    ['0xF6' CUSTOM_ERROR_118            ]
    ['0xF7' CUSTOM_ERROR_119            ]
    ['0xF8' CUSTOM_ERROR_120            ]
    ['0xF9' CUSTOM_ERROR_121            ]
    ['0xFA' CUSTOM_ERROR_122            ]
    ['0xFB' CUSTOM_ERROR_123            ]
    ['0xFC' CUSTOM_ERROR_124            ]
    ['0xFD' CUSTOM_ERROR_125            ]
    ['0xFE' CUSTOM_ERROR_126            ]
    ['0xFF' CUSTOM_ERROR_127            ]
]

[type HVACStatusFlags
    [simple   bit expansion     ]
    [simple   bit error         ]
    [simple   bit busy          ]
    [reserved bit 'false'       ]
    [simple   bit damperState   ]
    [virtual  bit isDamperStateClosed '!damperState']
    [virtual  bit isDamperStateOpen   'damperState' ]
    [simple   bit fanActive     ]
    [simple   bit heatingPlant  ]
    [simple   bit coolingPlant  ]
]

[type HVACHumidityModeAndFlags
    [reserved bit   'false'                                     ]
    [simple   bit   auxiliaryLevel                              ]
    [virtual  bit   isAuxLevelUnused '!auxiliaryLevel'          ]
    [virtual  bit   isAuxLevelUsed   'auxiliaryLevel'           ]
    [simple   bit   guard                                       ]
    [virtual  bit   isGuardDisabled '!guard'                    ]
    [virtual  bit   isGuardEnabled   'guard'                    ]
    [simple   bit   setback                                     ]
    [virtual  bit   isSetbackDisabled '!setback'                ]
    [virtual  bit   isSetbackEnabled   'setback'                ]
    [simple   bit   level                                       ]
    [virtual  bit   isLevelHumidity   '!level'                  ]
    [virtual  bit   isLevelRaw        'level'                   ]
    [simple   HVACHumidityModeAndFlagsMode  mode                ]
]

[enum uint 3 HVACHumidityModeAndFlagsMode
    ['0x0' OFF              ]
    ['0x1' HUMIDIFY_ONLY    ]
    ['0x2' DEHUMIDIFY_ONLY  ]
    ['0x3' HUMIDITY_CONTROL ]
]

[enum uint 8 HVACHumidityType
    ['0x00' NONE                        ]
    ['0x01' EVAPORATOR                  ]
    ['0x02' REFRIGERATIVE               ]
    ['0x03' EVAPORATOR_REFRIGERATIVE    ]
]

[enum uint 8 HVACHumidityError
    ['0x00' NO_ERROR                        ]
    ['0x01' HUMIDIFIER_TOTAL_FAILURE        ]
    ['0x02' DEHUMIDIFIER_TOTAL_FAILURE      ]
    ['0x03' FAN_TOTAL_FAILURE               ]
    ['0x04' HUMIDITY_SENSOR_FAILURE         ]
    ['0x05' HUMIDIFIER_TEMPORARY_PROBLEM    ]
    ['0x06' DEHUMIDIFIER_TEMPORARY_PROBLEM  ]
    ['0x07' FAN_TEMPORARY_PROBLEM           ]
    ['0x08' HUMIDIFIER_SERVICE_REQUIRED     ]
    ['0x09' DEHUMIDIFIER_SERVICE_REQUIRED   ]
    ['0x0A' FAN_SERVICE_REQUIRED            ]
    ['0x0B' FILTER_REPLACEMENT_REQUIRED     ]
    ['0x80' CUSTOM_ERROR_0                  ]
    ['0x81' CUSTOM_ERROR_1                  ]
    ['0x82' CUSTOM_ERROR_2                  ]
    ['0x83' CUSTOM_ERROR_3                  ]
    ['0x84' CUSTOM_ERROR_4                  ]
    ['0x85' CUSTOM_ERROR_5                  ]
    ['0x86' CUSTOM_ERROR_6                  ]
    ['0x87' CUSTOM_ERROR_7                  ]
    ['0x88' CUSTOM_ERROR_8                  ]
    ['0x89' CUSTOM_ERROR_9                  ]
    ['0x8A' CUSTOM_ERROR_10                 ]
    ['0x8B' CUSTOM_ERROR_11                 ]
    ['0x8C' CUSTOM_ERROR_12                 ]
    ['0x8D' CUSTOM_ERROR_13                 ]
    ['0x8E' CUSTOM_ERROR_14                 ]
    ['0x8F' CUSTOM_ERROR_15                 ]
    ['0x90' CUSTOM_ERROR_16                 ]
    ['0x91' CUSTOM_ERROR_17                 ]
    ['0x92' CUSTOM_ERROR_18                 ]
    ['0x93' CUSTOM_ERROR_19                 ]
    ['0x94' CUSTOM_ERROR_20                 ]
    ['0x95' CUSTOM_ERROR_21                 ]
    ['0x96' CUSTOM_ERROR_22                 ]
    ['0x97' CUSTOM_ERROR_23                 ]
    ['0x98' CUSTOM_ERROR_24                 ]
    ['0x99' CUSTOM_ERROR_25                 ]
    ['0x9A' CUSTOM_ERROR_26                 ]
    ['0x9B' CUSTOM_ERROR_27                 ]
    ['0x9C' CUSTOM_ERROR_28                 ]
    ['0x9D' CUSTOM_ERROR_29                 ]
    ['0x9E' CUSTOM_ERROR_30                 ]
    ['0x9F' CUSTOM_ERROR_31                 ]
    ['0xA0' CUSTOM_ERROR_32                 ]
    ['0xA1' CUSTOM_ERROR_33                 ]
    ['0xA2' CUSTOM_ERROR_34                 ]
    ['0xA3' CUSTOM_ERROR_35                 ]
    ['0xA4' CUSTOM_ERROR_36                 ]
    ['0xA5' CUSTOM_ERROR_37                 ]
    ['0xA6' CUSTOM_ERROR_38                 ]
    ['0xA7' CUSTOM_ERROR_39                 ]
    ['0xA8' CUSTOM_ERROR_40                 ]
    ['0xA9' CUSTOM_ERROR_41                 ]
    ['0xAA' CUSTOM_ERROR_42                 ]
    ['0xAB' CUSTOM_ERROR_43                 ]
    ['0xAC' CUSTOM_ERROR_44                 ]
    ['0xAD' CUSTOM_ERROR_45                 ]
    ['0xAE' CUSTOM_ERROR_46                 ]
    ['0xAF' CUSTOM_ERROR_47                 ]
    ['0xB0' CUSTOM_ERROR_48                 ]
    ['0xB1' CUSTOM_ERROR_49                 ]
    ['0xB2' CUSTOM_ERROR_50                 ]
    ['0xB3' CUSTOM_ERROR_51                 ]
    ['0xB4' CUSTOM_ERROR_52                 ]
    ['0xB5' CUSTOM_ERROR_53                 ]
    ['0xB6' CUSTOM_ERROR_54                 ]
    ['0xB7' CUSTOM_ERROR_55                 ]
    ['0xB8' CUSTOM_ERROR_56                 ]
    ['0xB9' CUSTOM_ERROR_57                 ]
    ['0xBA' CUSTOM_ERROR_58                 ]
    ['0xBB' CUSTOM_ERROR_59                 ]
    ['0xBC' CUSTOM_ERROR_60                 ]
    ['0xBD' CUSTOM_ERROR_61                 ]
    ['0xBE' CUSTOM_ERROR_62                 ]
    ['0xBF' CUSTOM_ERROR_63                 ]
    ['0xC0' CUSTOM_ERROR_64                 ]
    ['0xC1' CUSTOM_ERROR_65                 ]
    ['0xC2' CUSTOM_ERROR_66                 ]
    ['0xC3' CUSTOM_ERROR_67                 ]
    ['0xC4' CUSTOM_ERROR_68                 ]
    ['0xC5' CUSTOM_ERROR_69                 ]
    ['0xC6' CUSTOM_ERROR_70                 ]
    ['0xC7' CUSTOM_ERROR_71                 ]
    ['0xC8' CUSTOM_ERROR_72                 ]
    ['0xC9' CUSTOM_ERROR_73                 ]
    ['0xCA' CUSTOM_ERROR_74                 ]
    ['0xCB' CUSTOM_ERROR_75                 ]
    ['0xCC' CUSTOM_ERROR_76                 ]
    ['0xCD' CUSTOM_ERROR_77                 ]
    ['0xCE' CUSTOM_ERROR_78                 ]
    ['0xCF' CUSTOM_ERROR_79                 ]
    ['0xD0' CUSTOM_ERROR_80                 ]
    ['0xD1' CUSTOM_ERROR_81                 ]
    ['0xD2' CUSTOM_ERROR_82                 ]
    ['0xD3' CUSTOM_ERROR_83                 ]
    ['0xD4' CUSTOM_ERROR_84                 ]
    ['0xD5' CUSTOM_ERROR_85                 ]
    ['0xD6' CUSTOM_ERROR_86                 ]
    ['0xD7' CUSTOM_ERROR_87                 ]
    ['0xD8' CUSTOM_ERROR_88                 ]
    ['0xD9' CUSTOM_ERROR_89                 ]
    ['0xDA' CUSTOM_ERROR_90                 ]
    ['0xDB' CUSTOM_ERROR_91                 ]
    ['0xDC' CUSTOM_ERROR_92                 ]
    ['0xDD' CUSTOM_ERROR_93                 ]
    ['0xDE' CUSTOM_ERROR_94                 ]
    ['0xDF' CUSTOM_ERROR_95                 ]
    ['0xE0' CUSTOM_ERROR_96                 ]
    ['0xE1' CUSTOM_ERROR_97                 ]
    ['0xE2' CUSTOM_ERROR_98                 ]
    ['0xE3' CUSTOM_ERROR_99                 ]
    ['0xE4' CUSTOM_ERROR_100                ]
    ['0xE5' CUSTOM_ERROR_101                ]
    ['0xE6' CUSTOM_ERROR_102                ]
    ['0xE7' CUSTOM_ERROR_103                ]
    ['0xE8' CUSTOM_ERROR_104                ]
    ['0xE9' CUSTOM_ERROR_105                ]
    ['0xEA' CUSTOM_ERROR_106                ]
    ['0xEB' CUSTOM_ERROR_107                ]
    ['0xEC' CUSTOM_ERROR_108                ]
    ['0xED' CUSTOM_ERROR_109                ]
    ['0xEE' CUSTOM_ERROR_110                ]
    ['0xEF' CUSTOM_ERROR_111                ]
    ['0xF0' CUSTOM_ERROR_112                ]
    ['0xF1' CUSTOM_ERROR_113                ]
    ['0xF2' CUSTOM_ERROR_114                ]
    ['0xF3' CUSTOM_ERROR_115                ]
    ['0xF4' CUSTOM_ERROR_116                ]
    ['0xF5' CUSTOM_ERROR_117                ]
    ['0xF6' CUSTOM_ERROR_118                ]
    ['0xF7' CUSTOM_ERROR_119                ]
    ['0xF8' CUSTOM_ERROR_120                ]
    ['0xF9' CUSTOM_ERROR_121                ]
    ['0xFA' CUSTOM_ERROR_122                ]
    ['0xFB' CUSTOM_ERROR_123                ]
    ['0xFC' CUSTOM_ERROR_124                ]
    ['0xFD' CUSTOM_ERROR_125                ]
    ['0xFE' CUSTOM_ERROR_126                ]
    ['0xFF' CUSTOM_ERROR_127                ]
]

[type HVACHumidityStatusFlags
    [simple   bit expansion             ]
    [simple   bit error                 ]
    [simple   bit busy                  ]
    [reserved bit 'false'               ]
    [simple   bit damperState           ]
    [virtual  bit isDamperStateClosed '!damperState']
    [virtual  bit isDamperStateOpen   'damperState' ]
    [simple   bit fanActive             ]
    [simple   bit dehumidifyingPlant    ]
    [simple   bit humidifyingPlant      ]
]

[type HVACAuxiliaryLevel
    [reserved bit 'false'               ]
    [simple   bit fanMode               ]
    [virtual  bit isFanModeAutomatic  '!fanMode']
    [virtual  bit isFanModeContinuous 'fanMode' ]
    [simple   uint 6 mode               ]
    [virtual  bit isFanSpeedAtDefaultSpeed 'mode == 0x00'   ]
    [virtual  uint 6 speedSettings  'mode'                  ]
]

[enum uint 8 HVACSensorStatus
    ['0x00' NO_ERROR_OPERATING_NORMALLY                 ]
    ['0x01' SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND   ]
    ['0x02' SENSOR_OUT_OF_CALIBRATION                   ]
    ['0x03' SENSOR_TOTAL_FAILURE                        ]
]

[type HVACZoneList
    [simple  bit expansion              ]
    [simple  bit zone6                  ]
    [simple  bit zone5                  ]
    [simple  bit zone4                  ]
    [simple  bit zone3                  ]
    [simple  bit zone2                  ]
    [simple  bit zone1                  ]
    [simple  bit zone0                  ]
    [virtual bit unswitchedZone 'zone0' ]
]

[type HVACStartTime
    [simple  uint  16 minutesSinceSunday12AM                                    ]
    [virtual float 32 hoursSinceSunday12AM      'minutesSinceSunday12AM/60'     ]
    [virtual float 32 daysSinceSunday12AM       'hoursSinceSunday12AM/24'       ]
    [virtual uint  8  dayOfWeek                 'daysSinceSunday12AM+1'         ]
    [virtual uint  8  hour                      'hoursSinceSunday12AM%24'       ]
    [virtual uint  8  minute                    'minutesSinceSunday12AM%60'     ]
]
