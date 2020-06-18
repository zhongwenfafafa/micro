package main

import (
	"fmt"
	"log"

	"micro/bootstrap"
	"micro/router"
)

func main() {
	err := bootstrap.InitModule("./conf/dev")
	if err != nil {
		fmt.Println(err)
	}

	log.Fatal(router.Router().Run(":9999"))
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}
