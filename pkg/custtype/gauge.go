package custtype

import "encoding/json"

type gauge struct {
	value float64
}

func (c *gauge) GetVal() float64 {
	return c.value
}

func (c *gauge) Multi(anotherVal float64) CustomGauge {
	return &gauge{c.value * anotherVal}
}

func (c *gauge) Add(anotherVal float64) CustomGauge {
	return &gauge{c.value + anotherVal}
}

func (c *gauge) UnmarshalJSON(data []byte) error {
	custStruct := struct {
		Value float64 `json:"value"`
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
