package main

import (
	"fmt"
	"github.com/prometheus/common/log"
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
