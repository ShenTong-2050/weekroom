package main

import (
	"fmt"
)

type DumplingsShopInterface interface {
	Generate(t string) Dumplings
}

type Dumplings interface {
	create()
}

type BJMeatDumplings struct {}

func (*BJMeatDumplings) create() {
	fmt.Println("Beijing meat")
}

type BJVegetableDumplings struct {}

func (*BJVegetableDumplings) create() {
	fmt.Println("Beijing vegetable")
}

type SHMeatDumplings struct {}

func (*SHMeatDumplings) create() {
	fmt.Println("Shanghai meat")
}

type SHVegetableDumplings struct {}

func (*SHVegetableDumplings) create() {
	fmt.Println("Shanghai vegetable")
}

type BJDumplingsShop struct {}

type SHDumplingsShop struct {}

// Generate 北京工厂
func (*BJDumplingsShop) Generate(t string) Dumplings {
	switch t {
	case "meat":
		return new(BJMeatDumplings)
	case "vegetable":
		return new(BJVegetableDumplings)
	}
	return nil
}

// Generate 上海工厂
func (*SHDumplingsShop) Generate(t string) Dumplings {
	switch t {
	case "meat":
		return new(SHMeatDumplings)
	case "vegetable":
		return new(SHVegetableDumplings)
	}
	return nil
}

func main()  {

	/*fmt.Printf("%v,%x, %p\n",new(BJDumplingsShop),new(BJDumplingsShop))
	fmt.Printf("%v,%x, %p\n",BJDumplingsShop{},BJDumplingsShop{})*/

	var d Dumplings

	// 北京工厂
	BJFactory := new(BJDumplingsShop)
	d = BJFactory.Generate("meat")
	d.create()

	// 上海工厂
	// SHFactory := new(SHDumplingsShop)
	SHFactory := SHDumplingsShop{}
	d = SHFactory.Generate("vegetable")
	d.create()
}
