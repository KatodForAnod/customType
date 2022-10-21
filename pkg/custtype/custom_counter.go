package custtype

import (
	"encoding/json"
	"fmt"
	"sync"
)

type intTypeSwitcher string

const (
	int8Type      intTypeSwitcher = "int8"
	int16Type     intTypeSwitcher = "int16"
	int32Type     intTypeSwitcher = "int32"
	intCustomType intTypeSwitcher = "custom"

	maxInt32 = 2147483647
	maxInt16 = 32767
	maxInt8  = 127
)

type counterType struct {
	value     int64
	customMax int64

	switcher intTypeSwitcher
	sync.Mutex
}

func (s *counterType) initCounter(typeInt intTypeSwitcher, maxIntForCustom int64) error {
	switch typeInt {
	case int8Type:
		s.switcher = typeInt
	case int16Type:
		s.switcher = typeInt
	case int32Type:
		s.switcher = typeInt
	case intCustomType:
		s.switcher = typeInt
		s.customMax = maxIntForCustom
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
		return int32(s.value)
	}
	return 0
}

func (s *counterType) GetCounter8() int8 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == int8Type {
		return int8(s.value)
	}
	return 0
}

func (s *counterType) GetCounter16() int16 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == int16Type {
		return int16(s.value)
	}
	return 0
}

func (s *counterType) GetCounter() int64 {
	s.Lock()
	defer s.Unlock()

	if s.switcher == intCustomType {
		return s.value
	}
	return 0
}

func (s *counterType) GetMaxInt() int64 {
	return s.customMax
}

func (s *counterType) SetValue(value int64) error {
	err := fmt.Errorf("SetValue func err: value %d too huge for %s type", value, s.switcher)

	switch s.switcher {
	case int8Type:
		if value < maxInt8 {
			return err
		}
	case int16Type:
		if s.value < maxInt16 {
			return err
		}
	case int32Type:
		if s.value < maxInt32 {
			return err
		}
	case intCustomType:
		if s.value < s.customMax {
			return err
		}
	default:
		return fmt.Errorf("unsupport type %s", s.switcher)
	}

	s.value = value
	return nil
}

func (s *counterType) Increment() {
	s.Lock()
	defer s.Unlock()

	switch s.switcher {
	case int8Type:
		if s.value < maxInt8 {
			s.value = s.value + 1
		}
	case int16Type:
		if s.value < maxInt16 {
			s.value = s.value + 1
		}
	case int32Type:
		if s.value < maxInt32 {
			s.value = s.value + 1
		}
	case intCustomType:
		if s.value < s.customMax {
			s.value = s.value + 1
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

	err = s.SetValue(custStruct.Value)
	if err != nil {
		return fmt.Errorf("package: %s, func: %s, err: %s",
			"custtype", "UnmarshalJSON", err)
	}

	return nil
}
