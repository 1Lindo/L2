package visitorPkg

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}
