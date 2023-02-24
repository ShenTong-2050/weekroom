package main

import (
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

type resourcePoolConfig struct {
	name   string
	maxTotal,maxIdle,minIdle int
}

type resourcePoolConfigBuilder struct {
	name   string
	maxTotal,maxIdle,minIdle int
}

func (b *resourcePoolConfigBuilder) setName(name string) error {
	if name == "" {
		return fmt.Errorf("builder name %v cannot empty",name)
	}
	b.name = name
	return nil
}

func (b *resourcePoolConfigBuilder) setMaxTotal(total int) error {
	if total < 0 {
		return fmt.Errorf("builder total %v cannot < 0",total)
	}
	b.maxTotal = total
	return nil
}

func (b *resourcePoolConfigBuilder) setMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("builder maxIdle %v cannot < 0",maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *resourcePoolConfigBuilder) setMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("builder minIdle %v cannot < 0",minIdle)
	}
	b.minIdle = minIdle
	return nil
}

func (b *resourcePoolConfigBuilder) Build() (*resourcePoolConfig,error) {
	if b.name == "" {
		return nil,fmt.Errorf("builder name %v cannot empty",b.name)
	}
	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
		// return nil,fmt.Errorf("builder maxTotal %v cannnot equal 0",b.maxTotal)
	}
	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
		// return nil,fmt.Errorf("builder maxIdle %v cannot equal 0",b.maxIdle)
	}
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
		// return nil,fmt.Errorf("builder minIdle %v cannot equal 0",b.minIdle)
	}
	if b.maxIdle > b.maxTotal {
		return nil,fmt.Errorf("builder maxIdle %v cannot > maxTotal %v",b.maxIdle,b.maxTotal)
	}
	if b.minIdle > b.maxIdle {
		return nil,fmt.Errorf("builder minIdle %v cannot < maxIdle %v",b.minIdle,b.maxIdle)
	}
	return &resourcePoolConfig{
		name: b.name,
		maxTotal: b.maxTotal,
		maxIdle: b.maxIdle,
		minIdle: b.minIdle,
	},nil
}

func main() {

	var builder = resourcePoolConfigBuilder{name: "FirstRequest",maxTotal: defaultMaxTotal,maxIdle: 13,minIdle: 9}

	builder.setName("FirstRequest")

	got,err := builder.Build()

	fmt.Println(got)
	fmt.Println(err != nil)

}
