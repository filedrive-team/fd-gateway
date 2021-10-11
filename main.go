package main

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	cli "github.com/urfave/cli/v2"
)

func main() {

	logging.SetLogLevel("*", "debug")
	fmt.Println("开始")

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",
			Value: "~/.lotus",
		},
	}
	fmt.Println(app)
	fmt.Println("结束")
}
