package toolkit

import "fmt"

func Version() {
	fmt.Println("toolkit version: 0.0.1 ")
}

// If 三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
