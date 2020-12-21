package main

import "fmt"

func main(){
	nums := []int{2,3,4}
	sum := 0
	// 遍历数组 第一个元素为下标，第二个元素时下标上的元素
	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}

	// 遍历 kv,第一个元素时key，第二个元素时value
	kvs := map[string]string {"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}

	// 遍历字符串，第一个元素时下标，第二个元素时对应字符的unicode编码
	for i,c := range "go" {
		fmt.Println(i,c)
	}

}

/**
sum: 9
index: 1
a -> apple
b -> banana
0 103
1 111
 */
