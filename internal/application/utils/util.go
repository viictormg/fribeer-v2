package utils

import (
	"encoding/json"
	"fmt"
)

func PrintLog(data interface{}) {
	a, _ := json.Marshal(data)
	fmt.Println(string(a))
}
