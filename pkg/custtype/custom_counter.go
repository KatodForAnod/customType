package custtype

import (
	"encoding/json"
	"fmt"
	"sync"
)

type intTypeSwitcher string

const (
	int8Type  intTypeSwitcher = "int8"
	int16Type intTypeSwitcher = "int16"
	int32Type intTypeSwitcher = "int32"

	maxInt32 = 2147483647
	maxInt16 = 32767
	maxInt8  = 127
)

type counterType struct {
	value    interface{}
	switcher intTypeSwitcher
	sync.Mutex
}

func (s *counterType) initCounter(typeInt intTypeSwitcher) error {
	switch typeInt {
	case int8Type:
		s.switcher = typeInt
		s.value = int(0)
	case int16Type:
		s.switcher = typeInt
		s.value = int16(0)
	case int32Type:
		s.switcher = typeInt
		s.value = int32(0)
	default:
		return fmt.Errorf("package: %s, func: %s, err: %s",
			"custtype", "initCounter", "unsupport int type")
	}

	return nil
}

func (s *counterType) GetCounter32() int32 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == int32Type {
		if val, ok := s.value.(int32); ok {
			return val
		}
	}
	return 0
}

func (s *counterType) GetCounter8() int8 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == int8Type {
		if val, ok := s.value.(int8); ok {
			return val
		}
	}
	return 0
}

func (s *counterType) GetCounter16() int16 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == int16Type {
		if val, ok := s.value.(int16); ok {
			return val
		}
	}
	return 0
}

func (s *counterType) Increment() {
	s.Lock()
	defer s.Unlock()

	switch s.switcher {
	case int8Type:
		if s.value.(int8) < maxInt8 {
			s.value = s.value.(int8) + 1
		}
	case int16Type:
		if s.value.(int16) < maxInt16 {
			s.value = s.value.(int16) + 1
		}
	case int32Type:
		if s.value.(int32) < maxInt32 {
			s.value = s.value.(int32) + 1
		}
	default:
		return
	}
}

func (s *counterType) UnmarshalJSON(data []byte) error {
	custStruct := struct {
		Value int64 `json:"value"`
	}{
		Value: 0,
	}

	err := json.Unmarshal(data, &custStruct)
	if err != nil {
		return err
	}

	err = fmt.Errorf("package: %s, func: %s, err: %s",
		"custtype", "UnmarshalJSON", "unsupport int type")
	switch s.switcher {
	case int8Type:
		if custStruct.Value > maxInt8 {
			return err
		}
		s.value = int8(custStruct.Value)
	case int16Type:
		if custStruct.Value > maxInt16 {
			return err
		}
		s.value = int16(custStruct.Value)
	case int32Type:
		if custStruct.Value > maxInt32 {
			return err
		}
		s.value = int32(custStruct.Value)
	default:
		return err
	}

	return nil
}
