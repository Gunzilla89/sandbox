package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min int, max int) int {
	rand := min + rand.Intn(max-min)

	return rand
}

func randomString(len int) string {
	letterByte := make([]byte, len)

	for i := 0; i < len; i++ {
		letterByte[i] = byte(randomInt(60, 90))
	}

	return string(letterByte)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(randomString(25))
}
