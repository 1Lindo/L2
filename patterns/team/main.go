package main

import "main/pkg/teamPkg"

func main() {
	tv := &teamPkg.Tv{}

	onCommand := &teamPkg.OnCommand{
		Device: tv,
	}

	offCommand := &teamPkg.OffCommand{
		Device: tv,
	}

	onButton := &teamPkg.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &teamPkg.Button{
		Command: offCommand,
	}
	offButton.Press()
}
