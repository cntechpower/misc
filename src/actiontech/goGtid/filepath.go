package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	d := filepath.Dir("/tmp/test")
	fmt.Println(d)
}
