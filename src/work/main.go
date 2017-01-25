package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"redis_util"
)

func main() {
	fmt.Println("hello word")
	r := gin.Default()
	r.GET("/api/test", func(c *gin.Context) {
		redisClient := redis_util.GetRedisClientInstance()
		test := redisClient.Get("test")
		fmt.Println(test)
		if test == nil {
			c.JSON(200, gin.H{})
		} else {
			c.JSON(200, test)
		}
	})
	r.Run(":8080")
}
