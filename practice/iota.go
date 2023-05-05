package practice

import "fmt"

type FlagU uint8

const (
	FNone FlagU = iota
	FRead
	FWrite
	FRW      = FRead | FWrite
	FExecute = iota
	FPP
)

func F(f FlagU) {
	switch f {
	case FNone:
		fmt.Println("FNone")
	case FRead:
		fmt.Println("FRead")
	case FWrite:
		fmt.Println("FWrite")
	case FExecute:
		fmt.Println("FExecute")
	case FRW:
		fmt.Println("FRW")
	}
}
