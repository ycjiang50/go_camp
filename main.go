package main

import (
	"fmt"
	"go_camp/genericslice"
)

func main() {
	res := genericslice.Delete[int](2, []int{2, 3, 4, 5})
	fmt.Printf("new slice: %v, len: %d, cap: %d\n", res, len(res), cap(res))
	res = genericslice.Delete[int](2, res)
	fmt.Printf("new slice: %v, len: %d, cap: %d\n", res, len(res), cap(res))
}
