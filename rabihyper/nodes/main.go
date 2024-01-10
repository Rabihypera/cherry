package main

import (
	"fmt"
	rabihyperConst "github.com/rabihyper/internal/const"
	"github.com/rabihyper/nodes/gate"
	"github.com/rabihyper/nodes/master"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name:        "rabihyper game cluster node",
		Description: "rabihyper game cluster node release",
		Commands: []*cli.Command{
			versionCommand(),
			masterCommand(),
			gateCommand(),
		},
	}

	_ = app.Run(os.Args)
}

func versionCommand() *cli.Command {
	return &cli.Command{
		Name:      "version",
		Aliases:   []string{"ver", "v"},
		Usage:     "view version",
		UsageText: "rabihyper game cluster node version",
		Action: func(c *cli.Context) error {
			fmt.Println(rabihyperConst.Version())
			return nil
		},
	}
}

func masterCommand() *cli.Command {
	return &cli.Command{
		Name:      "master",
		Usage:     "run gate node",
		UsageText: "node master --path=./rabihyper/config/profile-rabihyper-development.json --node=rabihyper-master",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			master.Run(path, node)
			return nil
		},
	}
}

func gateCommand() *cli.Command {
	return &cli.Command{
		Name:      "gate",
		Usage:     "run gate node",
		UsageText: "node gate --path=./rabihyper/config/profile-rabihyper-development.json --node=rabihyper-gate-1",
		Flags:     getFlag(),
		Action: func(c *cli.Context) error {
			path, node := getParameters(c)
			gate.Run(path, node)
			return nil
		},
	}
}

func getParameters(c *cli.Context) (path, node string) {
	path = c.String("path")
	node = c.String("node")
	return path, node
}

func getFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Usage:    "profile config path",
			Required: false,
			Value:    "./rabihyper/config/profile-rabihyper-development.json",
		},
		&cli.StringFlag{
			Name:     "node",
			Usage:    "node id name",
			Required: true,
			Value:    "",
		},
	}
}
