package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/tobi007/angular-go-serve/bind"
	"github.com/tobi007/angular-go-serve/config"
	"github.com/tobi007/angular-go-serve/db"
	"github.com/tobi007/angular-go-serve/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	fs := &assetfs.AssetFS {
		Asset :      Asset ,
		AssetDir :   AssetDir ,
		AssetInfo : AssetInfo ,
		Prefix: "dist",
	}
	bfs := bind.NewBinaryFileSystem(fs)
	server.Init(bfs)
}
