package accounts

import (
	"fmt"
	"math/rand"
	"time"
)

func SignIn(email, number string) {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for x, y := range email {
		h := x + r1.Intn(5)
		fmt.Println(string(int(y)+h), y, h)
	}
}
