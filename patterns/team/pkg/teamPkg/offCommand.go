package teamPkg

type OffCommand struct {
	Device Device
}

func (c *OffCommand) execute() {
	c.Device.off()
}
