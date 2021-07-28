package utils

import (
	"math/rand"
	"time"
)

//随机生成一个长度为10的字符串
func RandomString(n int) string {
	var letters = []byte("hdfjigduysvhdfsbhjdsfgyudguavbajxiofheruibcfd")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix()) //随机数种子
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
