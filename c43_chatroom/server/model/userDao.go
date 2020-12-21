package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

// 与 redis 交互
type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{pool:pool}
	return
}

func (this *UserDao)getUserById(redisConn redis.Conn,id int)(user *User,err error){
	defer redisConn.Close()
	reply, err := redis.String(redisConn.Do("ping"))
	fmt.Println(reply)
	userInfo, err := redis.String(redisConn.Do("hget", "user", id))
	if err != nil {
		// 未找到
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
			return
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(userInfo),user)

	if err != nil {
		fmt.Printf("json.Unmarshal error: %v\n",err)
		return
	}else{
		fmt.Println("user from redis",*user)
	}
	return
}

func (this *UserDao)Login(userId int,userPwd string)(user *User,err error){
	redisConn := this.pool.Get()
	defer redisConn.Close()
	user, err = this.getUserById(redisConn, userId)
	if err != nil {
		return
	}

	// 校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

