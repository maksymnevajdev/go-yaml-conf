package main

import (
	"fmt"
	conf "github.com/maksymnevajdev/go-yaml-conf"
)

func main() {
	fmt.Println(conf.Local.GetUint("id"))
	fmt.Println(conf.Local.GetString("token_Secret"))
	fmt.Println(conf.Local.GetSlice("slice"))
}
