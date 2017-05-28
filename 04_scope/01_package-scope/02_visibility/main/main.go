package main

import (
	"fmt"
	"github.com/newveteran.07/01_go-training/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
