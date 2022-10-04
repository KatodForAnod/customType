package custtype

import (
	"encoding/json"
	"fmt"
)

type Builder struct {}

func (b *Builder) CreateCustomBool(jsonBytes []byte) (CustomBool, error) {
	newStruct := boolType{}
	err := json.Unmarshal(jsonBytes, &newStruct)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomBool err:%v", err)
	}

	return newStruct, nil
}

func (b *Builder) CreateCustomCounter8(jsonBytes []byte) (CustomCounter8, error) {
	newStruct := counterType{}
	newStruct.switcher = int8Type

	err := json.Unmarshal(jsonBytes, &newStruct)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter8 err:%v", err)
	}

	return &newStruct, nil
}

func (b *Builder) CreateCustomCounter16(jsonBytes []byte) (CustomCounter16, error) {
	newStruct := counterType{}
	newStruct.switcher = int16Type

	err := json.Unmarshal(jsonBytes, &newStruct)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter16 err:%v", err)
	}

	return &newStruct, nil
}

func (b *Builder) CreateCustomCounter32(jsonBytes []byte) (CustomCounter32, error) {
	newStruct := counterType{}
	newStruct.switcher = int32Type

	err := json.Unmarshal(jsonBytes, &newStruct)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomCounter32 err:%v", err)
	}

	return &newStruct, nil
}