package server

import (
	"fmt"
	"net/rpc"
	"os"
	"os/signal"

	"github.com/ilovelili/neo-go/config"
	"github.com/urfave/cli"
)

var (
	configPath = "./config"
)

// NewCommand creates a new Node command.
func NewCommand() cli.Command {
	return cli.Command{
		Name:   "node",
		Usage:  "start a NEO node",
		Action: serve,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "config-path"},
			cli.BoolFlag{Name: "privnet, p"},
			cli.BoolFlag{Name: "mainnet, m"},
			cli.BoolFlag{Name: "testnet, t"},
			cli.BoolFlag{Name: "debug, d"},
		},
	}
}

// serve start server
func serve(ctx *cli.Context) {
	net := resolveNetMode(ctx)

	configPath = ctx.String("config-path")
	cfg, err := config.Load(configPath, net)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	interruptChan := make(chan os.Signal)
	signal.Notify(interruptChan, os.Interrupt)

	serverConfig := network.NewServerConfig(cfg)
	chain, err := newBlockchain(cfg)

	if err != nil {
		err = fmt.Errorf("could not initialize blockchain: %s", err)
		return cli.NewExitError(err, 1)
	}

	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	server := network.NewServer(serverConfig, chain)
	rpcServer := rpc.NewServer(chain, cfg.ApplicationConfiguration.RPCPort, server)
	errChan := make(chan error)

	go server.Start(errChan)
	go rpcServer.Start(errChan)

	fmt.Println()
	fmt.Println(server.UserAgent)
	fmt.Println()
}

func resolveNetMode(ctx *cli.Context) config.NetMode {
	net := config.ModePrivNet
	if ctx.Bool("testnet") {
		net = config.ModeTestNet
	}
	if ctx.Bool("mainnet") {
		net = config.ModeMainNet
	}

	return net
}
