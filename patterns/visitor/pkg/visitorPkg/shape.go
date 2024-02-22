package visitorPkg

type Shape interface {
	getType() string
	accept(Visitor)
}
