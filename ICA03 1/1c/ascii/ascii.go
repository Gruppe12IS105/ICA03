// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package ascii

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func GreetingASCII() string {
	t1 := []byte("\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29\n")
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}
	t2 := "\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29"
	return t2

}
