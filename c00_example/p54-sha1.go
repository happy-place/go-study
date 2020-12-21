package main

import (
	"crypto/sha1"
	"fmt"
)

/**
求指定字符串的sha1散列值
 */
func main(){
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))

	// Sum appends the current hash to b and returns the resulting slice.
	bs := h.Sum(nil)
	fmt.Println(s) // sha1 this string
	fmt.Printf("%x\n",bs)
}

/**
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
 */