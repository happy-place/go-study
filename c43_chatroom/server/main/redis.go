package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string,maxIdle int,maxActive int,idleTimeout time.Duration){
	pool = &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         maxIdle,
		MaxActive:       maxActive,
		IdleTimeout:     idleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial:            func()(redis.Conn,error){
			return redis.Dial("tcp",address)
		},
	}
}
