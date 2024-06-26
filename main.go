package main

import (
	"fmt"
	"log"
	"os"

	"github.com/colinso/lego/config"
	"github.com/colinso/lego/generator"
	"github.com/colinso/lego/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "lego",
		Usage: "Build a microservice piece by piece",
		Commands: []*cli.Command{
			{
				Name:      "generate",
				Aliases:   []string{"gen", "g"},
				ArgsUsage: "lego generate <config path> <project path>",
				Usage:     "Generate code from a config file",
				Action: func(cCtx *cli.Context) error {
					configPath := cCtx.Args().First()
					projectPath := cCtx.Args().Get(1)
					if configPath == "" || projectPath == "" {
						fmt.Println("Please provide a config path and a project path")
						return nil
					}
					config.ParseConfig(configPath, projectPath)
					err := generator.NewCodeGenerator().Generate()
					if err != nil {
						return err
					}
					appName := config.GetConfig().Name

					utils.RunCleanupShellCommands(appName, projectPath)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
