package custtype

import "encoding/json"

type gauge struct {
	value int64
}

func (c *gauge) GetVal() int64 {
	return c.value
}

func (c *gauge) Multi(anotherVal int64) CustomGauge {
	return &gauge{c.value * anotherVal}
}

func (c *gauge) Add(anotherVal int64) CustomGauge {
	return &gauge{c.value + anotherVal}
}

func (c *gauge) UnmarshalJSON(data []byte) error {
	custStruct := struct {
		Value int64 `json:"value"`
	}{
		Value: 0,
	}

	err := json.Unmarshal(data, &custStruct)
	if err != nil {
		return err
	}

	c.value = custStruct.Value
	return nil
}
