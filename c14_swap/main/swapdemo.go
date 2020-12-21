package main

import "fmt"

func swap1(a *int,b *int){
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swap2(a *int,b *int){
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main(){
	a,b := 1,2
	swap1(&a,&b)
	fmt.Println("a=",a,"b=",b)
	swap2(&a,&b)
	fmt.Println("a=",a,"b=",b)

}
