package custtype

import (
	"encoding/json"
	"fmt"
)

type boolType struct {
	value     bool
	customOn  string
	customOff string
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

func (s *boolType) SetValue(value bool) {
	s.value = value
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
		custStruct.Value == "Off" {
		s.value = false
	} else {
		return fmt.Errorf("package: %s, func: %s, err: %s",
			"custtype", "UnmarshalJSON", "wrong bool type")
	}

	/*if custStruct.Value == s.customOn &&
		custStruct.Value != "" {
		s.value = true
	} else if custStruct.Value == s.customOff &&
		custStruct.Value != "" {
		s.value = false
	}*/

	return nil
}

func (s *boolType) String() string {
	var output string
	if s.value {
		output = "ON"
	} else {
		output = "OFF"
	}

	if s.customOn != "" && s.value == true {
		output = s.customOn
	} else if s.customOff != "" && s.value == false {
		output = s.customOff
	}

	return fmt.Sprintf(output)
}
