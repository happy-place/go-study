package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringLen(){
	// 统一按utf-8编码处理，ascii和数字占一个字符，汉子占3个字符
	str := "ab23中国" // 2 + 2 + 3 * 3 = 10
	fmt.Println(len(str)) // 数组长度
}

func stringLoop(){
	str := "ab23中国"
	for i:=0;i<len(str);i++{ // 汉子和ascii数字混杂，出现乱码
		fmt.Printf("%c ",str[i]) // a b 2 3 ä ¸ ­ å
	}

	str1 := []rune(str) // 使用[]rune 封装，解决乱码问题
	for i:=0;i<len(str1);i++{
		fmt.Printf("%c ",str1[i]) // a b 2 3 中 国
	}

	// range遍历，第一个元素是下标
	for _,e := range str{
		fmt.Printf("%c ",e) // a b 2 3 中 国
	}
}

func stringToInt(){
	a := "12"
	b,err := strconv.Atoi(a) // 字符串转整数，是下面strconv.ParseInt 的缩写
	if err != nil{
		panic(err) // 转换失败异常处理
	}
	fmt.Printf("b=%v, type: %T\n",b,b) // b=12, type: int

	c,_ := strconv.ParseInt(a,10,64)
	fmt.Printf("c=%v, type: %T\n",c,c)
}

func intToString(){
	a := 12
	b := strconv.Itoa(a)
	fmt.Printf("b=%v, type: %T\n",b,b) // b=12, type: string
}

func stringToByteArr(){
	// go没有char类型，使用byte表示char，即ascii码值代表字符
	a := []byte("hello")
	fmt.Println(a) // [104 101 108 108 111]
}

func byteArrToString(){
	a := string([]byte{104 ,101, 108, 108, 111})
	fmt.Printf("a=%s, type: %T\n",a,a)
}

func intToOther(){
	var a int64 = 123
	b := strconv.FormatInt(a, 2)
	o := strconv.FormatInt(a, 8)
	x := strconv.FormatInt(a, 16)

	fmt.Printf("b=%v, type: %T\n",b,b)
	fmt.Printf("b=%v, type: %T\n",o,o)
	fmt.Printf("b=%v, type: %T\n",x,x)
	/**
	b=1111011, type: string
	b=173, type: string
	b=7b, type: string

	*/
}

func strContains(){
	a := "hello"
	fmt.Println(strings.Contains(a,"he")) // true
	fmt.Println(strings.Contains(a,"eh")) // false
}

func strCount(){
	a := "hello"
	fmt.Println(strings.Count(a,"l")) // 2
}

func stringEqualIgnore(){
	fmt.Println("Abc"=="abc") // false 区分大小写
	fmt.Println(strings.EqualFold("aBc","ABC")) // true 不区分大小写
}

func indexOf(){
	// 首次匹配位置
	fmt.Println(strings.Index("hello","l"))
}

func lastIndexOf(){
	fmt.Println(strings.LastIndex("hello","l")) // 3 最后一个缩影
}

func relpace(){
	fmt.Println(strings.Replace("abababa","a","A",-1)) // AbAbAbA 替换全部
	fmt.Println(strings.Replace("abababa","a","A",2)) // 替换2次
}

func split(){
	print := func (arr []string){
		for _,e := range arr {
			fmt.Printf("'%s',",e)
		}
		fmt.Println()
	}
	arr1 := strings.Split("hello,hi,hey",",")
	print(arr1) // 'hello','hi','hey', 以指定分隔符分隔

	arr2 := strings.SplitAfter("hello,hi,hey",",")
	print(arr2) // 'hello,','hi,','hey', 从指定分隔符之后开始分隔
}

func toXXXCase(){
	fmt.Println(strings.ToUpper("abc")) // 转大写
	fmt.Println(strings.ToLower("ABC")) // 转小写
	fmt.Println(strings.ToTitle("Tom")) // 转标题
}

func trim(){
	fmt.Println(strings.Trim(" a b c  "," ")) // 掐头去尾
	fmt.Println(strings.TrimLeft(" a b c  "," ")) // 删左
	fmt.Println(strings.TrimRight(" a b c  "," ")) // 删右
}

func hasPrefix(){
	fmt.Println(strings.HasSuffix("hello","lo")) // 是否包含后缀
	fmt.Println(strings.HasPrefix("hello","he")) // 是否包含前缀
}

func main(){
	//stringLen()
	//stringLoop()
	//stringToInt()
	//intToString()
	//stringToByteArr()
	//byteArrToString()
	//intToOther()
	//strContains()
	//strCount()
	//stringEqualIgnore()
	//indexOf()
	//lastIndexOf()
	//relpace()
	//split()
	//toXXXCase()
	//trim()
	hasPrefix()
}


