package visitorPkg

type Rectangle struct {
	Side int
}

func (r *Rectangle) Accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}
