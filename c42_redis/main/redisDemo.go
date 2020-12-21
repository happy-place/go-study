package main

// go get github.com/garyburd/redigo/redis
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func String(){
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %s\n",err)
	}
	defer dial.Close()
	do, _ := dial.Do("set", "name", "tom")
	fmt.Printf("set name: %v\n",do) // set name: OK

	// 获取方式1
	nameBytes, _ := dial.Do("get", "name")
	fmt.Printf("get name: %v\n",string(nameBytes.([]byte))) // get name: tom

	name, _ := redis.String(dial.Do("get", "name"))
	fmt.Printf("get name: %v\n",name) // get name: tom

	do, _ = dial.Do("set", "age", 21)

	fmt.Printf("set age: %v\n",do) // set name: OK

	age, _ := redis.Int(dial.Do("get", "age"))
	fmt.Printf("age: %v, type: %T\n",age,age) // age: 21, type: int

}

func SetExpired(){
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %v\n",err)
	}
	defer dial.Close()
	dial.Do("set","name",3,"tom")
	time.Sleep(time.Second*3)
	name, err := dial.Do("get", "name")
	if err != nil {
		fmt.Printf("get name err: %v\n",err)
	}
	fmt.Printf("name is %v %v\n",nil,nil == name) // name is <nil> true

	dial.Do("set","name","tom")
	name, _ = redis.String(dial.Do("get", "name"))
	fmt.Printf("name: %s\n",name) // name: tom

	dial.Do("expire","name",3)

	time.Sleep(time.Second*3)
	name, err = dial.Do("get", "name")
	if err != nil {
		fmt.Printf("get name err: %v\n",err)
	}
	fmt.Printf("name: %v\n",name) // name: <nil>
}

func Hash(){
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %v\n",err)
	}
	defer dial.Close()
	dial.Do("hset","stu1","name","tom")
	dial.Do("hset","stu1","age",21)
	dial.Do("hset","stu1","address","Hubei-Wuhan")

	name, _ := redis.String(dial.Do("hget", "stu1", "name"))
	age,_ :=redis.Int(dial.Do("hget","stu1","age"))
	address,_ :=redis.String(dial.Do("hget","stu1","address"))
	fmt.Printf("name: %v,age: %v,address: %v\n",name,age,address)

	dial.Do("hmset","stu2","name","jack","address","Hubei-Macheng","age",23)
	// 类型明确，直接使用指定类型接收
	strings, _ := redis.Strings(dial.Do("hmget", "stu2", "name", "address"))
	for i,str := range strings {
		fmt.Printf("strings[%d] = %v\n",i,str)
	}

	// 类型混杂使用 []interface{} 接收
	values, err := redis.Values(dial.Do("hmget", "stu2", "name", "age", "address"))
	for i,str := range values {
		fmt.Printf("values[%d] = %v\n",i,string(str.([]byte)))
	}
}

type Monster struct {
	Id int
	Name string
	Age int
	Skil string
}

func WriteToRedis(m Monster){
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %v\n",err)
	}
	defer dial.Close()
	_, err = dial.Do("hmset", fmt.Sprintf("monster-%v", m.Id), "id", m.Id, "name", m.Name, "age", m.Age, "skill", m.Skil)
	if err != nil {
		fmt.Printf("save err: %v\n",err)
	}
}

func  ReadFromRedis(id int) {
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %v\n",err)
	}
	defer dial.Close()
	vv, _ := redis.Strings(dial.Do("hmget", fmt.Sprintf("monster-%v",id), "id", "name", "age", "skill"))
	//id,_ := strconv.Atoi(vv[0])
	name := vv[1]
	age,_ := strconv.Atoi(vv[2])
	skill := vv[3]
	monster := Monster{id, name, age, skill}
	fmt.Println(monster)
}

func GetAllFromRedis(){
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("connect to redis err: %v\n",err)
	}
	defer dial.Close()
	values, err := redis.Values(dial.Do("keys", "monster-*"))
	for _,v := range values {
		vv, _ := redis.Strings(dial.Do("hmget", string(v.([]byte)), "id", "name", "age", "skill"))
		id,_ := strconv.Atoi(vv[0])
		name := vv[1]
		age,_ := strconv.Atoi(vv[2])
		skill := vv[3]
		monster := Monster{id, name, age, skill}
		fmt.Println(monster)
	}
}

var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         8, // 最大空闲连接数
		MaxActive:       0, // 与数据库建立最大连接数，0表示不限制
		IdleTimeout:     300, // 最大空闲时间
		Wait:            false, // 达到最大连接数时，获取连接需要等待
		MaxConnLifetime: 0,
		Dial:  			func() (redis.Conn, error){
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

// 使用连接池避免反复创建连接开销
func TestPool(){
	defer pool.Close()
	conn := pool.Get()
	do, err := conn.Do("ping")
	if err != nil {
		fmt.Println("redis ping err: %v\n",err)
	}
	fmt.Printf("ping reply: %v\n",do) // ping reply: PONG
}

func MonsterTest(){
	//var name string
	//var age int
	//var skill string
	//for i:=0;i<3;i++{
	//	fmt.Printf("请输入第%v个monster的Name：",i+1)
	//	fmt.Scanf("%s",&name)
	//	fmt.Printf("请输入第%v个monster的Age：",i+1)
	//	fmt.Scanf("%d",&age)
	//	fmt.Printf("请输入第%v个monster的Skill：",i+1)
	//	fmt.Scanf("%s",&skill)
	//	monster := Monster{i, name, age, skill}
	//	WriteToRedis(monster)
	//}

	ReadFromRedis(1)
	//GetAllFromRedis()
}




func main(){
	//String()
	//SetExpired()
	//Hash()
	//MonsterTest()
	TestPool()
}
