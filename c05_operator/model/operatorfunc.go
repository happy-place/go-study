package model

import "fmt"

func GetWeeksAndDays(num int) string {
	weeks := num / 7
	days := num - weeks * 7
	if days == 0 {
		return fmt.Sprintf("%v个星期",weeks)
	}else {
		return fmt.Sprintf("%v个星期零%v天",weeks,days)
	}
}

func GetTemperature(huashi float32) float32 {
	return 5.0 / 9 * (huashi - 100)
}

func Swap(i int,j int) (int ,int){
	i = i ^ j
	j = i ^ j
	i = i ^ j
	return i,j
}

func Max(a ... int) int {
	max := a[0]
	for _,e := range a {
		if max < e {
			max = e
		}
	}
	return max
}
