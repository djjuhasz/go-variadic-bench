package main

import (
	"fmt"
	"time"

	"github.com/djjuhasz/go-variadic-bench/strf"
)

func main() {
	name := "David"
	fmt.Println(strf.F("Hello %s, happy %s!", name, time.Now().Weekday()))
}
