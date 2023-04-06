package util

import (
	"fmt"
	"strconv"
)

const (
	MaxUINT32           = 1<<32 - 1
	DefaultUuidCntCache = 512
)

type UuidGenerator struct {
	Prefix       string
	idGen        uint32
	internalChan chan uint32
}

func NewUuidGenerator(prefix string) *UuidGenerator {
	gen := &UuidGenerator{
		prefix,
		0,
		make(chan uint32, DefaultUuidCntCache),
	}
	gen.startGen()
	return gen
}

func (g *UuidGenerator) startGen() {
	go func() {
		for {
			if g.idGen == MaxUINT32 {
				g.idGen = 1
			} else {
				g.idGen += 1
			}
			g.internalChan <- g.idGen
		}
	}()
}

func (g *UuidGenerator) Get() string {
	idGen := <-g.internalChan
	return fmt.Sprintf("%s%d", g.Prefix, idGen)
}

func (g *UuidGenerator) GetUint32() string {
	idGen := <-g.internalChan
	maxLen := len(strconv.Itoa(MaxUINT32))
	fmtStr := fmt.Sprintf("0%dd", maxLen)
	return fmt.Sprintf("%"+fmtStr, idGen)
}
