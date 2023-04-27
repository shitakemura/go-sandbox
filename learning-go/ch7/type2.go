package main

import (
	"fmt"
	"io"
)

func doTypeSwitch(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("nil", j)
	case int:
		fmt.Println("int", j)
	case string:
		fmt.Println("string", j)
	default:
		fmt.Println("other", j)
	}
}

func main() {
	doTypeSwitch(10)
	doTypeSwitch("hoge")
	doTypeSwitch(nil)
	doTypeSwitch(true)
}

func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (writter int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}

	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	return
}
