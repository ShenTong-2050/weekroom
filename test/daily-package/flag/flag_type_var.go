package main

import (
	"flag"
	"fmt"
)

var (
	intFlag 	int
	boolFlag	bool
	stringFlag	string
)

func init() {
	flag.IntVar(&intFlag,"intFlag",1,"intFlag description")
	flag.BoolVar(&boolFlag,"boolFlag",true,"boolFlag description")
	flag.StringVar(&stringFlag,"stringFlag","default","stringFlag description")
}

func main() {
	flag.Parse()
	fmt.Println("the intFlag is:",intFlag)
	fmt.Println("the boolFlag is:",boolFlag)
	fmt.Println("the stringFlag is:",stringFlag)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	flag.PrintDefaults()
}
