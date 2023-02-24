package main

type Nike struct {}

type NikeShoes struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (s Nike) MakeShoes() MyShoes {
	return &NikeShoes{Shoe{logo: "nike",size: 12}}
}

func (s Nike) MakeShirt() MyShirt {
	return &NikeShirt{Shirt{logo: "nike",size: 13}}
}
