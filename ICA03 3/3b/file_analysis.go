package main

import (
	"fmt"

	"./fileutils"
)

func main() {
	filename := "./files/lang01.wl"

	var f = fileutils.FileToByteslice(filename)
	fmt.Printf("% X", f)
}
