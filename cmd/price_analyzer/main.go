package main

import (
	"fmt"
	"price_analyzer_prototype/internal/app"
	"price_analyzer_prototype/internal/price_analyzer_api"
)

var (
	// filled with LDFLAGS
	version, service string
	_, _             = version, service
)

func main() {
	var cfg = app.Config{
		Server: struct {
			Host string
			Port int
		}{
			"localhost",
			8080,
		},
		Database: struct {
			User   string
			Pass   string
			Host   string
			Port   int
			Dbname string
		}{
			User:   "postgres",
			Pass:   "password",
			Host:   "localhost",
			Port:   5432,
			Dbname: "price_analyze",
		},
	}

	proxyAddr := ":8080"
	serviceAddr := "127.0.0.1:8081"

	dbConn, err := app.NewDB(cfg)
	if err != nil {
		fmt.Println("error connect to database: %w", err)
		return
	}

	priceAnalyzeServer := price_analyzer_api.NewProductAnalyzerAPI(dbConn)
	go app.GRPCService(serviceAddr, priceAnalyzeServer)
	app.HTTPProxy(proxyAddr, serviceAddr)
	//c := core.New(version, service, &cfg)
	//c.Run(app.New(c, cfg))
}
