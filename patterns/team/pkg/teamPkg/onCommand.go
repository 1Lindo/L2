package teamPkg

type OnCommand struct {
	Device Device
}

func (c *OnCommand) execute() {
	c.Device.on()
}
