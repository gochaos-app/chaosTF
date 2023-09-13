package main

import (
	"log"
	"os"

	"github.com/gochaos-app/chaosTF/exec"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "chaosTF",
		Usage: "A terminal based app that transforms your infrastructure code into go-chaos config files",
		Commands: []*cli.Command{
			{
				Name: "readtf",
				Usage: `Read terraform file: 
				chaosTF readtf main.tf env:dev
				chaosTF readtf main.tf env:dev myFile.hcl`,
				Action: func(cCtx *cli.Context) error {
					tfFile := cCtx.Args().Get(0)
					pattern := cCtx.Args().Get(1)
					chaosFile := cCtx.Args().Get(2)
					key, value := exec.TagCheck(pattern)
					if chaosFile == "" {
						chaosFile = "chaos-config.hcl"
					}
					exec.ExecConfig(tfFile, key, value, chaosFile)
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
