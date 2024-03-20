package main

import (
	"log"
	"os"

	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/generator"
	"github.com/colinso/lego/actions/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lego",
		Usage: "Build a microservice piece by piece",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "Generate code from a config file",
				Action: func(cCtx *cli.Context) error {
					configPath := cCtx.Args().First()
					projectPath := cCtx.Args().Get(1)
					config.ParseConfig(configPath, projectPath)
					err := generator.NewCodeGenerator().Generate()
					if err != nil {
						return err
					}
					appName := config.GetConfig().Name
					//TODO: Update this to take location as an argument
					utils.RunCleanupShellCommands(appName, projectPath+"/"+appName)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
