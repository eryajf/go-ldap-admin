package tools

import (
	"fmt"
	"testing"
)

func TestGenPass(t *testing.T) {
	fmt.Printf("密码为：%s\n", NewGenPasswd("123456"))
	// err := ComparePasswd("$2a$10$Fy8p0nCixgWKzLfO3SgdhOzAF7YolSt6dHj6QidDGYlzLJDpniXB6", "123456")
	// if err != nil {
	// 	fmt.Printf("密码错误：%s\n", err)
	// }
}

func TestArrUintCmp(t *testing.T) {
	a := []uint{1, 2, 3, 4, 6, 9}
	b := []uint{1, 2, 3, 4, 5, 6, 7}
	c, d := ArrUintCmp(a, b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}

func TestSliceToString(t *testing.T) {
	a := []uint{1}
	fmt.Printf("%s\n", SliceToString(a, ","))
}
