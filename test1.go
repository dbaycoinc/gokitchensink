package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "evangeline@educative.io"
	x := strings.Split(email, "@")
	z := strings.TrimLeft()
	y := (strings.Replace("XXX", x[0], "", 5)) + "@" + x[1]
	fmt.Println(y)
}
