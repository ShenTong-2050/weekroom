package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	intFlags		*int
	boolFlags		*bool
	stringFlags		*string
	period			time.Duration
)

func init() {
	intFlags = flag.Int("intFlag",88,"intFlag description")
	boolFlags = flag.Bool("boolFlag",false,"boolFlag description")
	stringFlags = flag.String("stringFlag","default string flag","string description")
	flag.DurationVar(&period,"period",1,"default sleep time")
	// period = flag.Duration("period",1,"sleep describe")
}

func main() {
	flag.Parse()

	time.Sleep(period)
	fmt.Printf("sleeping for %v\n",period)

	fmt.Println("intFlags is: ",*intFlags)
	fmt.Println("boolFlags is: ",*boolFlags)
	fmt.Println("stringFlags is: ",*stringFlags)
}
