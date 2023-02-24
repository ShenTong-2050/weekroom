package main

type MySports interface {
	MakeShoes() MyShoes
	MakeShirt() MyShirt
}

func GetMySports (brand string) MySports {
	switch brand {
	case "adidas":
		return &Adidas{}
	case "nike":
		return &Nike{}
	}
	return nil
}


