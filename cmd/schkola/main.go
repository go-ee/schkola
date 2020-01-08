package main

import (
	"os"

	"github.com/go-ee/schkola/app"
	"github.com/go-ee/utils/eh/app/memory"
	"github.com/go-ee/utils/eh/app/mongo"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli"
)

var Log = lg.NewLogger("CGT ")

func main() {

	var flag_secure = "secure"
	var flag_name = "name"

	name := "Schkola"

	runner := cli.NewApp()
	runner.Usage = name
	runner.Version = "1.0"
	runner.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "name, n",
			Usage: "name of the scool, used for backend data",
			Value: "schkola",
		},
		&cli.BoolFlag{
			Name:  "secure, s",
			Usage: "activate secure mode",
		},
	}

	flag_url := "url"
	runner.Commands = []*cli.Command{
		{
			Name:  "mongo",
			Usage: "Start server with MongoDB backend",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  flag_url,
					Usage: "url of the MongoDB instance",
					Value: "localhost",
				},
			},
			Action: func(c *cli.Context) (err error) {
				schkola := app.NewSchkola(mongo.NewAppMongo(name, c.String(flag_name), c.GlobalBool(flag_secure), c.String(flag_url)))
				schkola.Start()
				return
			},
		}, {
			Name:  "memory",
			Usage: "Start server with memory backend",
			Action: func(c *cli.Context) (err error) {
				schkola := app.NewSchkola(memory.NewAppMemory(name, c.String(flag_name), c.GlobalBool(flag_secure)))
				schkola.Start()
				return
			},
		},
	}

	if err := runner.Run(os.Args); err != nil {
		Log.Err("%v", err)
	}
}
