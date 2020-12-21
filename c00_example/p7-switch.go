package main

import (
	"fmt"
	"time"
)

func main(){
	// 标准switch case 进行值匹配
	i := 2
	fmt.Print("write ",1," as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 枚举类型匹配
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a weekday")
	}

	// case 中进行大小匹配
	t := time.Now()
	switch {
		case t.Hour() < 12 :
			fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}

	// 接收任意类型对象，然后使用 i.(type)进行类型匹配
	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("I`m a bool")
		case int:
			fmt.Println("I`m an int")
		default:
			fmt.Printf("Don`t know type %T\n",t) // 只有 Printf 才能正常打印 %T
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
/*
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
*/

