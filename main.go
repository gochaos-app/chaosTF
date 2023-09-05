package main

import (
	"log"
	"os"

	"github.com/gochaos-app/chaosTF/exec"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mentalTF",
		Usage: "A terminal based app that transforms your infrastructure code into go-chaos config files",
		Commands: []*cli.Command{
			{
				Name:  "readtf",
				Usage: "Read terraform file",
				Action: func(cCtx *cli.Context) error {
					tfFile := cCtx.Args().Get(0)

					pattern := cCtx.Args().Get(1)

					exec.ExecConfig(tfFile, pattern)
					return nil
				},
			},
			{
				Name:  "readstate",
				Usage: "Read terraform file",
				Action: func(cCtx *cli.Context) error {
					tfFile := cCtx.Args().Get(0)
					patternTag := cCtx.Args().Get(1)

					exec.ExecConfig(tfFile, patternTag)
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
