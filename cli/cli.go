package cli

import (
	"log"
	"os"

	"github.com/gochaos-app/chaosTF/exec"
	"github.com/urfave/cli/v2"
)

func ExecCli() {
	app := &cli.App{
		Name:  "chaosTF",
		Usage: "A terminal based app that transforms your infrastructure code into go-chaos config files",
		Commands: []*cli.Command{
			{
				Name: "read",
				Usage: `Reads a terraform file:
				chaos read main.tf env:dev
				chaosTF readtf main.tf env:dev basic`,
				Action: func(c *cli.Context) error {
					tfFile := c.Args().Get(0)
					tag := c.Args().Get(1)
					action := c.Args().Get(2)
					exec.LogicSingleFileRead(tfFile, tag, action)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
