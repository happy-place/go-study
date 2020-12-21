package main

import (
	"fmt"
	"reflect"
)

// 1.参数为空接口，即可以接受任意类型
func basicType(i interface{}){
	rType := reflect.TypeOf(i)  // 获取reflect.Type类型，如果是结构体，可以直接调用其绑定的方法
	rTypeName := rType.Name()  // 获取reflect.Type类型的名称，对应的就是原始类型名称
	rValue := reflect.ValueOf(i) // 获取reflect.Value类型，如果是基本数据类型，可以获取具体值
	fmt.Printf("type=%T name=%s value=%v\n",rType,rTypeName,rValue)

	var n1 int64
	n1 = rValue.Int() + 1 // 获取值方法1：显示获取值
	fmt.Println("n1=",n1)

	num := rValue.Interface() // 获取值方法2：先将rValue转为空接口类型，然后对接口使用类型断言，转为原始类型对象
	i2 := num.(int)
	var n2 int = i2
	fmt.Println("n2=",n2)
}

type Student struct{
	Name string
	Age int
}

type Fruit struct {
	Name string
	Color string
}

func structType(i interface{}){
	rType := reflect.TypeOf(i) // reflect.Type 指向入参的真实类型
	tKind := rType.Kind() // reflect.Type 对象类别（struct int float ...）
	rValue := reflect.ValueOf(i) // reflect.Value 的值
	vkind := rValue.Kind() // reflect.Value的类别 与 rType.Kind() 相同
	fmt.Printf("rType:%v tKind:%v rValue:%v vkind:%v\n",rType,tKind,rValue,vkind)

	// 获取原类型的对象
	// 方法1：类型断言 + bool 判断是否断言成功
	i2 := rValue.Interface()
	student,ok := i2.(Student)
	if ok {
		fmt.Println(student)
	}

	// 方法2：switch-case中类型断言，枚举处理
	switch i2.(type) {
	case Student:{
		s := i2.(Student)
		fmt.Printf("name:%s, age:%d\n",s.Name,s.Age)
	}
	case Fruit:
		f := i2.(Fruit)
		fmt.Printf("name:%s, color:%s\n",f.Name,f.Color)
	default:
		fmt.Println("未知类型！")
	}

}

// 传入指针，通过反射修改指针指向地址的内容
func changeBasicValue(i interface{}){
	rTypeOf := reflect.TypeOf(i)
	fmt.Printf("type:%v kind:%v\n",rTypeOf,rTypeOf.Kind())
	rValue := reflect.ValueOf(i)
	// Elem returns the value that the interface v contains 返回空接口指向的对象，或指针指向的空间
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	rValue.Elem().SetInt(20)
	// 等效于
	//n := 10
	//var b *int = &n
	//*b = 20
}

func test1(){
	f := 1.0
	rType := reflect.TypeOf(f)
	fmt.Printf("type:%s, kind:%s\n",rType,rType.Kind()) // type:float64, kind:float64

	rValue := reflect.ValueOf(f)
	fmt.Println(rValue.Float())

	i := rValue.Interface()
	f2 := i.(float64)
	fmt.Println(f2)
}

func test2(){
	s1 := "hello"
	rValue := reflect.ValueOf(&s1) // 反射修改值，必须基于指针
	rValue.Elem().SetString("hi")
	fmt.Println(s1)
}



// 基于反射修改结构体类型对象（传地址）
func test3(){
	stu := Student{"tom",21}
	rType := reflect.TypeOf(&stu)
	fmt.Printf("type: %s, kind:%s\n",rType,rType.Kind()) // type: *main.Student, kind:ptr

	// 基于反射给修改属性
	rValue := reflect.ValueOf(&stu)

	// 方法1：重新拿到元素结构体对象，以点的方式修改
	i := rValue.Interface()
	stu2 := i.(*Student) // 类型断言 获取结构体指针
	stu2.Name = "jack"
	fmt.Println(stu) // {jack 21}

	// 直接基于反射Setter赋值
	rValue.Elem().FieldByName("Name").SetString("tom")
	fmt.Println(rValue)

}

type Dog struct {
	Name string  `json:"name"`
	Age int `json:"age"`
}

func (dog Dog)Shut(){
	fmt.Printf("%s is wang wang~\n",dog.Name)
}

func (dog Dog)Eat(food string){
	fmt.Printf("%s is eating %s\n",dog.Name,food)
}

func TestStruct(data interface{}){
	rType := reflect.TypeOf(data)
	rValue := reflect.ValueOf(data)
	kind := rType.Kind()
	if kind != reflect.Struct {
		// 只解析结构体
		return
	}

	// 访问结构体属性
	// 方法1：基于下标访问（属性名称ascii排序）
	numOfFields := rValue.NumField() // 只能访问公有属性（大写）
	for i:=0;i<numOfFields;i++ {
		field := rType.Field(i)
		fmt.Printf("name: %v, type: %v, index: %v, pkgPath:%v\n",field.Name,field.Type,field.Index,field.PkgPath)
		jsonTagVal := field.Tag.Get("json")
		if jsonTagVal != "" {
			fmt.Printf("json-tag: %v\n",jsonTagVal)
		}
	}

	// 访问2：基于属性名称访问
	nameField, found := rType.FieldByName("Name")
	if found{
		fmt.Printf("name: %v, type: %v, index: %v, pkgPath:%v\n",nameField.Name,nameField.Type,nameField.Index,nameField.PkgPath)
		jsonTagVal := nameField.Tag.Get("json") // 获取结构体标签
		if jsonTagVal != "" {
			fmt.Printf("json-tag: %v\n",jsonTagVal)
		}
	}

	numOfMethods := rValue.NumMethod() // 只能访问公有方法（大写）
	fmt.Printf("struct %s has %d method\n",rType,numOfMethods)

	// 方法调用
	// 方法1： rValue.MethodByName 直接基于方法名调用
	// 调用dog.Shut()
	shutFunc := rValue.MethodByName("Shut")
	shutFunc.Call(nil)

	// 调用dog.Eat()
	eatFunc := rValue.MethodByName("Eat")
	var params []reflect.Value
	params = append(params, reflect.ValueOf("meat"))
	eatFunc.Call(params)

	// 方法2：rValue.Method(0) 通过下标获取方法（按名称排序）
	eatFunc2 := rValue.Method(0)
	eatFunc2.Call(params)

	shutFunc2 := rValue.Method(1)
	shutFunc2.Call(nil)

}

func test4(){
	dog := Dog{"tom",3}
	TestStruct(dog)
}

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal)GetSub()int {
	return cal.Num1 - cal.Num2
}

func test5(){
	var cal = &Cal{} // 此处必须是实质存对象指针 不能是 val cal *Cal
	rType := reflect.TypeOf(cal)
	rValue := reflect.ValueOf(cal)
	numOfFields := rType.Elem().NumField()
	for i:=0;i<numOfFields;i++{
		field := rType.Elem().Field(i)
		fmt.Printf("name: %v, type: %v, tag:%v\n",field.Name,field.Type,field.Tag)
	}
	obj := rValue.Elem()
	obj.FieldByName("Num1").SetInt(1)
	obj.FieldByName("Num2").SetInt(2)
	result := obj.MethodByName("GetSub").Call(nil)
	i := result[0].Int()
	fmt.Printf("result=%v\n",i)
}

func main(){
	//n := 100
	//basicType(n)

	//s := Student{"tom",21}
	//fmt.Println(s)
	//p := Fruit{"apple","green"}
	//structType(p)

	//num := 10
	//changeBasicValue(&num)
	//fmt.Println(num)

	//test1()
	//test2()
	test3()

	//test4()
	//test5()

}
