package model

import "fmt"

func PrintBin(nums ... int){
	for _,e := range nums {
		fmt.Printf("%b\n",e)
	}
}
