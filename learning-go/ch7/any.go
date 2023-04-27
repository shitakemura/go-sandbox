package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := map[string]any{}
	contents, err := os.ReadFile("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(contents, &data)
	fmt.Println(data)
}
