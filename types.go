package types

// Float is a constraint that permits any floating type number
type Float interface {
	~float64 | ~float32
}

// Signed is a constraint that permits any signed integer
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Int is a constraint that permits any integer
type Int interface {
	Signed | Unsigned
}

// Number is a constraint that permits any number
type Number interface {
	Float | Int
}
