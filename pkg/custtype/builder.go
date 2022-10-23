package custtype

import (
	"encoding/json"
	"fmt"
	"time"
)

func CreateCustomBool(jsonBytes []byte, customOn, customOff string) (CustomBool, error) {
	newStruct := boolType{
		customOff: customOff,
		customOn:  customOn,
	}
	err := json.Unmarshal(jsonBytes, &newStruct)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomBool err:%v", err)
	}

	return &newStruct, nil
}

func CreateCustomCounter8() (CustomCounter8, error) {
	newStruct := counterType{}
	err := newStruct.initCounter(int8Type, 0)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter8 err: %v", err)
	}

	return &newStruct, nil
}

func CreateCustomCounter16() (CustomCounter16, error) {
	newStruct := counterType{}
	err := newStruct.initCounter(int16Type, 0)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter16 err: %v", err)
	}

	return &newStruct, nil
}

func CreateCustomCounter32() (CustomCounter32, error) {
	newStruct := counterType{}
	err := newStruct.initCounter(int32Type, 0)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter32 err: %v", err)
	}

	return &newStruct, nil
}

func CreateCustomCounter(maxVal int64) (CustomCounter, error) {
	newStruct := counterType{}
	err := newStruct.initCounter(intCustomType, maxVal)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter err: %v", err)
	}

	return &newStruct, nil
}

func CreateCustomGauge(value int64) (CustomGauge, error) {
	newStruct := gauge{value: value}
	return &newStruct, nil
}

func CreateCustomTimestamp(tm time.Time) (CustomTimestamp, error) {
	newStruct := timeStamp{value: tm}
	return &newStruct, nil
}
