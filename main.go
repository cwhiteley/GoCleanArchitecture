package main

import (
	"os"

	"recipes/appcontext"
	"recipes/console"
	"recipes/server"

	"github.com/urfave/cli"
)

func main() {
	appcontext.Initialize()
	clientApp := cli.NewApp()
	clientApp.Name = "recipes-app"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start an HTTP API server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run database migrations",
			Action: func(c *cli.Context) error {
				return console.RunDatabaseMigrations()
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				return console.RollbackLatestMigration()
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
