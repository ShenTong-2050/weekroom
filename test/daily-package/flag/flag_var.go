package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// 自定义类型 interval
type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval has declared")
	}
	for _,dt := range strings.Split(value,"-") {
		duration,err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i,duration)
	}
	return nil
}

var (
	intervalFlag interval
)

func init() {
	flag.Var(&intervalFlag,"deltaT","custom arg deltaT description")
}

func main() {
	flag.Parse()
	fmt.Println(intervalFlag)
}
