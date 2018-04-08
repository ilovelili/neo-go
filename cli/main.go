package main

import (
	"os"

	"github.com/ilovelili/neo-go/cli/server"
	"github.com/ilovelili/neo-go/cli/smartcontract"
	"github.com/ilovelili/neo-go/cli/vm"
	"github.com/ilovelili/neo-go/cli/wallet"
	"github.com/urfave/cli"
)

func main() {
	client := cli.NewApp()
	client.Name = "neo-go"
	client.Usage = "Go client for Neo."

	client.Commands = []cli.Command{
		server.NewCommand(),
		smartcontract.NewCommand(),
		wallet.NewCommand(),
		vm.NewCommand(),
	}

	client.Run(os.Args)
}
