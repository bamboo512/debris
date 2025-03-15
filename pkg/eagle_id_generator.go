package pkg

import (
	"math/rand"
	"time"
)

const (
	idLength = 13
	charSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateEagleRandomID(prefix string) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为种子

	// 确保prefix 长度小于 idLength
	if len(prefix) > idLength {
		prefix = prefix[:idLength]
	}

	// 创建一个长度为 idLength 的 byte slice, 使用prefix填充
	id := make([]byte, idLength)
	for i := range prefix {
		id[i] = prefix[i]
	}

	// 填充剩余的字符
	for i := len(prefix); i < idLength; i++ {
		id[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(id)
}
