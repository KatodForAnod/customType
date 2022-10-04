package custtype

type CustomString interface {
	
}

type Gauge interface {

}

type CustomBool interface {

}

type CustomCounter8 interface {
	GetCounter8() int8
	Increment()
}

type CustomCounter16 interface {
	GetCounter16() int16
	Increment()
}

type CustomCounter32 interface {
	GetCounter32() int32
	Increment()
}

type CustomTimestamp interface {

}