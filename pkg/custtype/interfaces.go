package custtype

import "time"

type CustomString interface {
}

type CustomGauge interface {
	GetVal() int64
	Multi(anotherVal int64) CustomGauge
	Add(anotherVal int64) CustomGauge
}

type CustomBool interface {
	Off()
	On()
	GetBool() bool
	SetValue(value bool)
}

type CustomCounter8 interface {
	GetCounter8() int8
	Increment()
	SetValue(value int64) error
}

type CustomCounter16 interface {
	GetCounter16() int16
	Increment()
	SetValue(value int64) error
}

type CustomCounter32 interface {
	GetCounter32() int32
	Increment()
	SetValue(value int64) error
}

type CustomCounter interface {
	GetCounter() int64
	Increment()
	SetValue(value int64) error
}

type CustomTimestamp interface {
	GetValue() time.Time
	TimeNowDifference() time.Duration
}
