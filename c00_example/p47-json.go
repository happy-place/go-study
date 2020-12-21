package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int
	Fruits []string
}

func main(){
	// bool 类型序列化
	bolB,_ := json.Marshal(true)
	fmt.Println(bolB,string(bolB)) // [116 114 117 101] true

	intB,_ := json.Marshal(1)
	fmt.Println(intB,string(intB)) // [49] 1

	fltB,_ := json.Marshal(2.34)
	fmt.Println(fltB,string(fltB)) // [50 46 51 52] 2.34

	strB,_ := json.Marshal("gopher")
	fmt.Println(strB,string(strB)) // [34 103 111 112 104 101 114 34] "gopher" 包括引号

	slcD := []string{"apple","peach","pear"}
	slcB,_ := json.Marshal(slcD)
	fmt.Println(slcB,string(slcB)) // [91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93] ["apple","peach","pear"]

	mapD := map[string]int{"apple":5,"lettuce":7}
	mapB,_ := json.Marshal(mapD)
	fmt.Println(mapB,string(mapB)) // [123 34 97 112 112 108 101 34 58 53 44 34 108 101 116 116 117 99 101 34 58 55 125] {"apple":5,"lettuce":7}

	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple","peach","pear"}}
	res1B,_ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := Response2{
		Page:1,
		Fruits: []string{"apple","peach","pear"}}
	res2B,_ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err:= json.Unmarshal(byt,&dat);err !=nil {
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]] 字符串反序列化成map对象

	num := dat["num"].(float64) // 取出值
	fmt.Println(num) // 6.13

	strs := dat["strs"].([]interface{}) // 取出数组
	str1 := strs[0].(string) // 数组元素按字符串处理
	fmt.Println(str1) // a

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str),&res) // 字节数组反序列化
	fmt.Println(res) // &{1 [apple peach]}
	fmt.Println(res.Fruits[0]) // apple

	// {"apple":5,"lettuce":7} 输出
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d)

}