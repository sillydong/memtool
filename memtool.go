package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/bradfitz/gomemcache/memcache"
	"fmt"
	"strconv"
	"time"
)

const APPVERSION = "20160622"

var app *cli.App
var client *memcache.Client

func main() {
	app = cli.NewApp()
	app.Name = "MemcacheTool"
	app.Author = "Chen.Zhidong"
	app.Copyright = "https://sillydong.com"
	app.Usage = "Easier way to communicate with memcache"
	app.Version = APPVERSION
	app.Commands = []cli.Command{
		{
			Name:"get",
			Usage:"get value of key",
			ArgsUsage:"key [key...]",
			Action:get,
			SkipFlagParsing:true,
		},
		{
			Name:"set",
			Usage:"set value for key",
			ArgsUsage:"key value expiration",
			Action:set,
			SkipFlagParsing:true,
		},
		{
			Name:"del",
			Usage:"delete key",
			ArgsUsage:"key [key...]",
			Action:del,
			SkipFlagParsing:true,
		},
		{
			Name:"flush",
			Usage:"flush all keys",
			ArgsUsage:" ",
			Action:flush,
			SkipFlagParsing:true,
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"host",
			Value:"127.0.0.1",
			Usage:"Host to connect",
		},
		cli.StringFlag{
			Name:"port",
			Value:"11211",
			Usage:"Port to connect",
		},
	}
	app.Run(os.Args)
}

func connect(ctx *cli.Context) error {
	host := ctx.GlobalString("host")
	port := ctx.GlobalString("port")
	if len(host) > 0 && len(port) > 0 {
		client = memcache.New(fmt.Sprintf("%v:%v", host, port))
		return nil
	}else {
		return fmt.Errorf("missing memcache host and port")
	}
}

func get(ctx *cli.Context) error {
	if err := connect(ctx); err != nil {
		return err
	}
	if ctx.NArg() == 0 {
		return cli.ShowCommandHelp(ctx, "get")
	}else {
		items, err := client.GetMulti(ctx.Args())
		if err != nil {
			return err
		}else {
			if len(items) > 0 {
				for key, item := range items {
					if item == nil {
						fmt.Println("%v: not exist\n", key)
					}else {
						fmt.Printf("%v: %v\n", item.Key, string(item.Value))
					}
				}
			}else {
				fmt.Println("nothing found")
			}
		}
	}

	return nil
}

func set(ctx *cli.Context) error {
	if err := connect(ctx); err != nil {
		return err
	}
	if ctx.NArg() != 2 && ctx.NArg() != 3 {
		return cli.ShowCommandHelp(ctx, "set")
	}else {
		args := ctx.Args()
		if (ctx.NArg() == 2) {
			item := &memcache.Item{}
			item.Key = args.Get(0)
			item.Value = []byte( args.Get(1))
			err := client.Set(item)
			if err != nil {
				return err
			}else {
				fmt.Println("success")
				return nil
			}
		}else {
			item := &memcache.Item{}
			item.Key = args.Get(0)
			item.Value = []byte( args.Get(1))
			expire, err := strconv.Atoi(args.Get(2))
			if err != nil {
				return nil
			}else {
				item.Expiration = int32(int(time.Now().Unix()) + expire)
				fmt.Println(item.Expiration)
				err := client.Set(item)
				if err != nil {
					return err
				}else {
					fmt.Println("success")
					return nil
				}
			}
		}
	}
	return nil
}

func del(ctx *cli.Context) error {
	if err := connect(ctx); err != nil {
		return err
	}

	if ctx.NArg() == 0 {
		return cli.ShowCommandHelp(ctx, "del")
	}else {
		args := ctx.Args()
		for _, arg := range args {
			fmt.Println("delete " + arg + " ...")
			err := client.Delete(arg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func flush(ctx *cli.Context) error {
	if err := connect(ctx); err != nil {
		return err
	}

	err := client.FlushAll()
	if err != nil {
		return err
	}
	return nil
}
