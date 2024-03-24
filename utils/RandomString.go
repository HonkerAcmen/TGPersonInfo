package utils

import (
	"math/rand"
	"time"
)

func RandomString(num int) string {
	var str = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	var result = make([]byte, num)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = str[rand.Intn(len(str))]
	}
	return string(result)
}
func RandomInt(n int) string {
	var str = []byte("0123456789")
	var result = make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = str[rand.Intn(len(str))]
	}
	return string(result)
}
