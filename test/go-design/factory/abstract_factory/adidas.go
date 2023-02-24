package main

type Adidas struct {}

type AdidasShoes struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a Adidas) MakeShoes() MyShoes {
	return &AdidasShoes{Shoe{logo: "adidas",size: 16}}
}

func (a Adidas) MakeShirt() MyShirt {
	return &AdidasShirt{Shirt{logo: "adidas",size: 14}}
}
