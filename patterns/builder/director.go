package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setDirector(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildDb() Db {
	d.builder.setConfig()
	d.builder.setConn()
	d.builder.pingDb()
	return d.builder.getDb()
}
