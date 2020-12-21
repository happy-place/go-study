package main

import "time"
import "fmt"
import "math/rand"

func main(){
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println() // 81,87

	fmt.Println(rand.Float64()) // 0.6645600532184904

	fmt.Print((rand.Float64()*5)+5,",")
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println() // 7.1885709359349015,7.123187485356329

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100),",")
	fmt.Print(r1.Intn(100))
	fmt.Println() // 80,0

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100),",")
	fmt.Print(r2.Intn(100))
	fmt.Println() // 5,87

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100),",")
	fmt.Print(r3.Intn(100)) // 5,87
}