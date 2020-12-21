package main

import (
	"fmt"
	"strconv"
)

func strtobasic(){
	var str string = "true"
	var b bool
	b ,_= strconv.ParseBool(str) // func ParseBool(str string) (bool, error)
	fmt.Printf("b=%v, b type:%T\n",b,b) // b=true, b type:bool

	var str2 string = "12345690"
	var i1 int64
	var i2 int
	i1,_ = strconv.ParseInt(str2,10,64)
	i2 = int(i1)
	fmt.Printf("i1=%v, i1 type: %T\n",i1,i1) // i1=12345690, i1 type: int64
	fmt.Printf("i2=%v, i2 type: %T\n",i2,i2) // i2=12345690, i2 type: int

	var str3 string = "123.456"
	var f1 float64
	var f2 float32
	f1,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("f1=%v, f1 type: %T\n",f1,f1) // f1=123.456, f1 type: float64

	// 强转
	f2 = float32(f1)
	fmt.Printf("f2=%v, f2 type: %T\n",f2,f2)

	// 类型错误导致转换失败，不处理异常情况下，得到的也是错误值（默认值）
	var str4 string = "hello"
	var num int64
	num,err := strconv.ParseInt(str4,10,64)
	if err != nil {
		panic(err) // strconv.ParseInt: parsing "hello": invalid syntax
	}
	fmt.Printf("num=%v, num type: %T\n",num,num)




}

func basictostring(){
	var num1 int = 99
	var num11 int64 = 99
	var num2 float64 = 23.456
	var num3 bool = true
	//var myChar byte = 'h'
	var str string

	// 方式2： strconv
	str = strconv.Itoa(num1) // 等效于 FormatInt(int，10)
	str = strconv.FormatInt(num11,10) // 对 int64进行转换
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="99",str type: string

	// 'f'（-ddd.dddd)，10: 小数位保留10位，64：float64
	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="23.4560000000",str type: string
	/**
	bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	*/

	str = strconv.FormatBool(num3)
	fmt.Printf("str=%q,str type: %T\n",str,str) // str="true",str type: string

}

func main(){
	basictostring()
	strtobasic()
}
