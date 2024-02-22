package chainPkg

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
