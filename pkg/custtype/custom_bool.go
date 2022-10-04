package custtype

import (
	"encoding/json"
	"fmt"
)

type boolType struct {
	value      bool
}

func (s *boolType) GetBool() bool {
	return s.value
}

func (s *boolType) On() {
	s.value = true
}

func (s *boolType) Off() {
	s.value = false
}

func (s *boolType) UnmarshalJSON(data []byte) error {
	custStruct := struct {
		Value string
	}{
		Value: "",
	}

	err := json.Unmarshal(data, &custStruct)
	if err != nil {
		return err
	}

	if custStruct.Value == "On" ||
		custStruct.Value == "ON" {
		s.value = true
	} else if custStruct.Value == "OFF" ||
		custStruct.Value == "Of" {
		s.value = false
	} else {
		return fmt.Errorf("package: %s, func: %s, err: %s",
			"custtype", "UnmarshalJSON", "wrong bool type")
	}

	return nil
}
