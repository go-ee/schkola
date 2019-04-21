package main

import (
	"os"

	"github.com/go-ee/schkola/person"
	"github.com/go-ee/utils/eio"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli"
)

var Log = lg.NewLogger("CGT ")

func main() {

	var flag_url = "url"
	var flag_count = "count"
	var flag_file = "file"

	name := "Schkola Creator"

	runner := cli.NewApp()
	runner.Usage = name
	runner.Version = "1.0"

	runner.Commands = []cli.Command{
		{
			Name:  "generateChurches",
			Usage: "Generate churches demo entities to JSON file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  flag_file,
					Usage: "Target file",
					Value: "churches.json",
				},
				cli.IntFlag{
					Name:  flag_count,
					Usage: "Count of church entities to build",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) (err error) {
				file := c.String(flag_file)
				count := c.Int(flag_count)

				churches := person.BuildChurches(count)

				eio.CreateFileJSON(churches, file)

				return
			},
		},
		{
			Name:  "createChurches",
			Usage: "Create churches demo entities over REST API",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  flag_url,
					Usage: "Url of the REST API",
					Value: "http://127.0.0.1:8080",
				},
				cli.IntFlag{
					Name:  flag_count,
					Usage: "Count of church entities to build",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) (err error) {
				url := c.String(flag_url)
				count := c.Int(flag_count)

				churches := person.BuildChurches(count)

				creator := person.NewChurchImporter(url)
				err = creator.Create(churches)
				return
			},
		},
		{
			Name:  "importChurches",
			Usage: "Import churches file over REST API",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  flag_url,
					Usage: "Url of the REST API",
					Value: "http://127.0.0.1:8080",
				},
				cli.StringFlag{
					Name:  flag_file,
					Usage: "Source JSON file",
					Value: "churches.json",
				},
			},
			Action: func(c *cli.Context) (err error) {
				url := c.String(flag_url)
				file := c.String(flag_file)

				creator := person.NewChurchImporter(url)
				err = creator.ImportJSON(file)
				return
			},
		},
	}

	if err := runner.Run(os.Args); err != nil {
		Log.Err("%v", err)
	}
}
