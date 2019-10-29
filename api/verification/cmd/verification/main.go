package main

import (
	"fmt"

	"github.com/utilitywarehouse/partner-bazel-hello/api/account/pkg/cuba"
	_ "github.com/utilitywarehouse/partner-pkg/operational"
)

func main() {
	e := NewEvent("123")
	e.Comment = "foo-bar"
	fmt.Printf("verification api, yep: %+v\n", e)
	fmt.Println(cuba.Name("Rob"))
}
