package main

import (
	"customType/pkg/custtype"
	"fmt"
)

func main() {
	builder := custtype.Builder{}
	jsonStr := `{"value":"OFF"}`

	counter32, err := builder.CreateCustomBool([]byte(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(counter32)
}
