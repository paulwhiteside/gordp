package functions

import (
	"fmt"
)

func Foo(arg ...interface{}) interface{} {

	var tot float64
	for i := range arg {
		v := arg[i]

		fmt.Println("=============================>", v)

		switch v.(type) {
		case int:
			tot += float64(v.(int))
		case float64:
			tot += v.(float64)
		}

	}

	fmt.Println("returning", tot)
	return tot
}
