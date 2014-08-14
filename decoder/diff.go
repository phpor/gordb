package decoder


import (
	"fmt"

	"github.com/phpor/gordb/parser"
)

type Diff struct {
	db int
	i  int
	parser.NopDecoder
}

func (p *Diff) StartDatabase(n int) {
	p.db = n
}

func (p *Diff) Set(key, value []byte, expiry int64) {
	fmt.Printf("db=%d %q -> %q\n", p.db, key, value)
}

func (p *Diff) Hset(key, field, value []byte) {
	fmt.Printf("db=%d %q . %q -> %q\n", p.db, key, field, value)
}

func (p *Diff) OnHset(key, field, value []byte) {}
func (p *Diff) Sadd(key, member []byte) {
	fmt.Printf("db=%d %q { %q }\n", p.db, key, member)
}

func (p *Diff) StartList(key []byte, length, expiry int64) {
	p.i = 0
}

func (p *Diff) Rpush(key, value []byte) {
	fmt.Printf("db=%d %q[%d] -> %q\n", p.db, key, p.i, value)
	p.i++
}

func (p *Diff) Zadd(key []byte, score float64, member []byte) {
	fmt.Printf("db=%d %q[%d] -> {%q, score=%g}\n", p.db, key, p.i, member, score)
	p.i++
}

func (p *Diff) StartZSet(key []byte, cardinality, expiry int64) {
	p.i = 0
}
