package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"go.xargs.dev/archy"
	"go.xargs.dev/archy/option"
	"go.xargs.dev/archy/source"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

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
