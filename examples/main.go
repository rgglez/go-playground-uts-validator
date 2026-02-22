package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	isuts "github.com/rgglez/go-playground-uts-validator"
)

type Payload struct {
	Timestamp string `validate:"required,uts"`
}

func main() {
	v := validator.New()
	if err := isuts.RegisterUTSValidator(v, ""); err != nil {
		panic(err)
	}

	ok := Payload{Timestamp: "1735689600"}
	bad := Payload{Timestamp: "not-a-timestamp"}

	fmt.Println(v.Struct(ok) == nil)  // true
	fmt.Println(v.Struct(bad) == nil) // false
}