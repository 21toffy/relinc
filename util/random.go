package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alaphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alaphabet)
	for i := 0; i < k; i++ {
		c := alaphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

	// rand.Seed(time.Now().UnixNano())
	// letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// b := make([]rune, n)
	// for i := range b {
	//     b[i] = letters[rand.Intn(len(letters))]
	// }
	// return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, NRA, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomAccountType() string {
	currencies := []string{"SAVINGS", "CURRENT", "KIDDIES", "COPORATE"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomPhone() string {
	prefixes := []string{"080", "081", "090", "091", "070"}
	n := len(prefixes)
	prefix := prefixes[rand.Intn(n)]
	var numbers string
	for i := 0; i < 8; i++ {
		randomNumber := rand.Intn(10)
		numbers += strconv.Itoa(randomNumber)
	}
	return prefix + numbers

}
