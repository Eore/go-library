package main

import (
	"encoding/json"
	"fmt"

	errorlib "github.com/Eore/go-library/error"
)

var err error = errorlib.NewError(errorlib.TypeInfo, "MIAW").WithMessage("miawmiaw")

func main() {
	e := err
	b, _ := json.Marshal(e)
	fmt.Println(e)
	fmt.Println(string(b))
}
