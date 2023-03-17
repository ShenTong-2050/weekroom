package main

import (
	"flag"
	"fmt"
)

func main() {

	var args = []string{"-intFlag","12","-stringFlag","test"}

	var (
		intFlag		int
		boolFlag	bool
		stringFlag	string
	)

	fs := flag.NewFlagSet("MyFlagSet",flag.ContinueOnError)

	fs.IntVar(&intFlag,"intFlag",1,"default intFlag description")
	fs.BoolVar(&boolFlag,"boolFlag",false,"default boolFlag description")
	fs.StringVar(&stringFlag,"stringFlag","default string","default stringFlag description")

	fs.Parse(args)

	fmt.Println("intFlag is : ",intFlag)
	fmt.Println("boolFlag is : ",boolFlag)
	fmt.Println("stringFlag is : ",stringFlag)
}
