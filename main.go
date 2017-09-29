package main

import (
	"fmt"
)

type output struct {
	str string
}

func (o *output) set() {
	o.str = "hello world!"
}

func (o *output) String() string {
	return fmt.Sprintf("%s", o.str)
}

func main() {
	o := new(output)
	o.set()
	fmt.Println(o)
}
