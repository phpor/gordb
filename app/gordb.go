// This is only for me
package main

import (
	"fmt"
	"os"

	"github.com/phpor/gordb/parser"
	rdbdecoder "github.com/phpor/gordb/decoder"
	"flag"
)
type zipmapbindecoder struct {
	rdbdecoder.Zipmapbin
	max uint64
	i uint64
	verbos bool
	vunit uint64
}
func (p *zipmapbindecoder) Hset(key, field, value []byte) {
	p.Zipmapbin.Hset(key, field, value)
	p.i++
	if p.verbos && p.vunit > 0 {
		showProgress(p.i, p.vunit)
	}
	if p.max > 0 && p.i >= p.max {
		os.Exit(2)
	}
}
type zipmapdiffdecoder struct {
	rdbdecoder.Diff
	max uint64
	i uint64
	verbos bool
	vunit uint64
}
func (p *zipmapdiffdecoder) Hset(key, field, value []byte) {
	p.Diff.Hset(key, field, value)
	p.i++
	if p.verbos && p.vunit > 0 {
		showProgress(p.i, p.vunit)
	}
	if p.max > 0 && p.i >= p.max {
		os.Exit(2)
	}
}

func showProgress(num, unit uint64) {
	if num % unit == 0 {
		fmt.Fprintf(os.Stderr, ".")
		if num / unit % 80 == 0 {
			fmt.Fprintf(os.Stderr, "\t%d\n", num)
		}
	}
}
func maybeFatal(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	format := flag.String("format", "diff bin", "out format")
	filename := flag.String("file", "", "rdb file name")
	num := flag.Uint64("num", uint64(0), "num to dump")
	verbos := flag.Bool("verbos", false, "show progress")
	vunit := flag.Uint64("vunit", uint64(10), "unit of progress")

	flag.Parse()

	var f *os.File
	if *filename == "" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(*filename)
		maybeFatal(err)
	}

	var d parser.Decoder
	switch *format{
	case "bin":
		d = &zipmapbindecoder{max:*num, verbos:*verbos, vunit:*vunit}
	case "diff":
		d = &zipmapdiffdecoder{max:*num, verbos:*verbos, vunit:*vunit}
	default:
		flag.Usage();os.Exit(1)
	}
	err := parser.Decode(f, d)
	maybeFatal(err)
}
