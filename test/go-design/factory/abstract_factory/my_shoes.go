package main

type MyShoes interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

type Shoe struct {
	logo  string
	size  int
}

func(s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}