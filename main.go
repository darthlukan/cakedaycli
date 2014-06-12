package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/darthlukan/cakeday"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "findcakeday"
	app.Usage = "Find your Reddit Cake Day"
	app.Action = func(c *cli.Context) {
		username := c.Args()[0]
		response, err := cakeday.Get(username)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", response)
	}
	app.Run(os.Args)
}
