package main

import "fmt"

func main(){
	fmt.Println("姓名\t性别\t籍贯\t住址\nTom\t男\t湖北\t武汉")
	fmt.Println("顺顺利利\r麻麻") // \r 回到行首，继续输入的"麻麻"覆盖掉了已经存在的"顺顺"
}

/**
姓名    性别    籍贯    住址
Tom     男      湖北    武汉

麻麻利利
 */


