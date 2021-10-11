package main

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ipfs/go-log/v2"
	cli "github.com/urfave/cli/v2"
)

func main() {

	logging.SetLogLevel("*", "info")

	local := []*cli.Command{
		getCmd,
		postCmd,
	}

	app := &cli.App{
		Name:     "fd_gateway",
		Flags:    []cli.Flag{},
		Commands: local,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(os.Args)
		fmt.Println(err)
	}
}

var getCmd = &cli.Command{
	Name:  "get",
	Usage: "Get files by CID",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		targetPath := c.Args().First()

		res, err := gateway.GetByCID(ctx, targetPath)

		if err != nil {
			return err
		}

		fmt.Println(res)

		return nil
	},
}

var postCmd = &cli.Command{
	Name:  "post",
	Usage: "Post files into the IPFS system",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		targetPath := c.Args().First()

		res, err := gateway.PostFiles(ctx, targetPath)

		if err != nil {
			return err
		}

		fmt.Println(res)

		return nil
	},
}
