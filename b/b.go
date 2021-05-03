package b

import (
	"a"
	"fmt"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	{
		fmt.Println("From()")
		a.FromA()
	}
}
