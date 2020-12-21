package model

import (
	"fmt"
	"os"
)

func doErr(err error){
	if err != nil{
		panic(err)
	}
}

func closeFile(file *os.File){
	_ = file.Close()
	fmt.Println("closeFile")
}

func Openfile(name string){
	file,err := os.Open(name)
	defer closeFile(file)
	doErr(err)
}
