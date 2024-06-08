package main

import (
	"encoding/json"
	"fmt"
	"mojang/majong"
)

func main() {
	majong.ThisGame.Start()
}

func PrettyPrint(v any) {
	a, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		fmt.Println(v)
		return
	}
	fmt.Println(string(a))
}
