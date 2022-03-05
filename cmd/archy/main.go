package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"

	"go.xargs.dev/archy"
	"go.xargs.dev/archy/option"
	"go.xargs.dev/archy/source"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if len(os.Args) >= 2 && os.Args[1] == "version" {
		version()
		return
	}

	o := option.ParseFlags()

	v := &archy.Values{}
	if err := source.Use(o.Source).Values(ctx, o, v); err != nil {
		log.Fatalf("reading values from %s: ", o.Source, err)
	}

	result := os.Stdout

	if o.JSON {
		if err := json.NewEncoder(result).Encode(v); err != nil {
			log.Fatalf("encoding result as json: ", err)
		}
	} else {
		result.Write([]byte(v.String()))
		result.Write([]byte("\n"))
	}
}

func version() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Printf("unable to retrieve build info")
		os.Exit(1)
	}
	fmt.Println(bi.Main.Version)
}
