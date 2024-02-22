package builderPkg

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetDirector(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildDb() Db {
	d.builder.SetConfig()
	d.builder.SetConn()
	d.builder.PingDb()
	return d.builder.GetDb()
}
