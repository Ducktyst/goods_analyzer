package main

import (
	"price_analyzer_prototype/internal/app"
)

var (
	// filled with LDFLAGS
	version, service string
	_, _             = version, service
)

func main() {
	var cfg app.Config
	c := core.New(version, service, &cfg)
	c.Run(app.New(c, cfg))
}
