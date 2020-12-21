package main

import (
	"encoding/json" // 可以对 map 切片 struct 和 基本数据类型进行序列化
	"fmt"
)

// 结构体标签
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Hobby []string `json:"hobby"`
	Score map[string]float64 `json:"score"`
}

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Hobby []string `json:"hobby"`
	Score map[string]float64 `json:"score"`
}


func sayHello(name string){
	fmt.Println("hello,",name)
}

// 序列化 与 反序列化，目标对象类型需要一致
// 对于结构体而言类型一致，指的是字段名称、类型、顺序完全一致，至于结构体名称可以不同
func testJson(){
	person := Person{"tom", 23, "male",make([]string,0),make(map[string]float64)}
	person.Hobby = append(person.Hobby, "swimming","skitting")
	person.Score["English"] = 90.0
	person.Score["Chinese"] = 95.8
	fmt.Printf("%+v\n",person)

	// func Marshal(v interface{}) ([]byte, error)
	personBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("person 序列化失败(%s)\n",err)
	}
	fmt.Println(string(personBytes))
	var person2 Person // 反序列化时，map 和 切片 都不需要make，Unmarshal内部进行了make
	json.Unmarshal(personBytes,&person2)
	fmt.Printf("%+v\n",person2)

	var stu Student
	json.Unmarshal(personBytes,&stu) // person字符串完美反序列化为Student
	fmt.Printf("%v\n",stu)

}

func testFunc(){
	marshal, err := json.Marshal(sayHello)
	if err != nil {
		fmt.Printf("sayHello函数序列化失败(%s)\n",err)
		// sayHello函数序列化失败(json: unsupported type: func(string))
	}
	ff := func(name string){}
	json.Unmarshal(marshal,&ff)
	ff("tom")
}

func testDeserialize(){
	person := Person{"tom", 23, "male",make([]string,0),make(map[string]float64)}
	person.Hobby = append(person.Hobby, "swimming","skitting")
	person.Score["English"] = 90.0
	person.Score["Chinese"] = 95.8

	personBytes, _ := json.Marshal(person)
	fmt.Println(string(personBytes))
	var p1 Person
	json.Unmarshal([]byte(string(personBytes)),&p1)
	fmt.Println(p1)

	// 手动将json字符串 反序列化为 对象
	var p2 Person
	str2 := "{\"name\":\"tom\",\"age\":23,\"gender\":\"male\",\"hobby\":[\"swimming\",\"skitting\"],\"score\":{\"Chinese\":95.8,\"English\":90}}"
	fmt.Println(str2)
	json.Unmarshal([]byte(str2),&p2)
	fmt.Println(p2)

}

func main(){
	//testJson()
	//testFunc()
	testDeserialize()
}

