package main

import "github.com/chhz0/goiam/internal/apisvr"

func main() {
	if err := apisvr.New().Execute(); err != nil {
		panic(err)
	}
}
