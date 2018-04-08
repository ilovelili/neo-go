package vm

import "github.com/urfave/cli"

// NewCommand creates a new VM command.
func NewCommand() cli.Command {
	return cli.Command{
		Name:   "vm",
		Usage:  "start the neo virtual machine",
		Action: startVMPrompt,
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "debug, d"},
		},
	}
}

func startVMPrompt(ctx *cli.Context) error {
	p := vmcli.New()
	return p.Run()
}
