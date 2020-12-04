package main

import (
	"fmt"

	"gitlab.pactindo.com/internet-banking-busines/product-ibb-2.0/ibb-2.0-services/services/ibb-admin-service/helper/mapper"
)

type Data1 struct {
	N string
	M int
}

type Data2 struct {
	N string
	M int
}

func main() {
	d1 := Data1{N: "Test", M: 2}
	d2 := new(Data2)
	if err := mapper.Map(d1, d2); err != nil {
		panic(err.Error())
	}

	fmt.Println(d2)
}
