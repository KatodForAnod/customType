package custtype

import "time"

type timeStamp struct {
	value time.Time
}

func (t *timeStamp) GetValue() time.Time {
	return t.value
}

func (t *timeStamp) TimeNowDifference() time.Duration {
	return time.Now().Sub(t.value)
}
