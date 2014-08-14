package decoder

import (
	"os"

	"github.com/phpor/gordb/parser"
	"encoding/binary"
)

type Zipmapbin struct {
	db int
	i  int
	parser.NopDecoder
}

func (p *Zipmapbin) StartDatabase(n int) {
	p.db = n
}

func (p *Zipmapbin) Hset(key, field, value []byte) {
	lenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lenByte, uint32(len(key)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(key)

	binary.BigEndian.PutUint32(lenByte, uint32(len(field)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(field)

	binary.BigEndian.PutUint32(lenByte, uint32(len(value)))
	os.Stdout.Write(lenByte)
	os.Stdout.Write(value)
}
