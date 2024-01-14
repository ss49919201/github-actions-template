package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b, _ := json.Marshal(map[string]any{"foo": "bar"})
	fmt.Println(string(b))
}
