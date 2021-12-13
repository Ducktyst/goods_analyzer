package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/jmoiron/sqlx"
	"log"
	"net"
	"price_analyzer_prototype/internal/api"
	"sync"
	"time"
)

type App struct {
	config Config

	productAnalyzerAPI api.ProductAnalyzerAPI
	server *Server
	DB     *pg.Options
}

type Server struct {
	Host string
	Port int
	mu   sync.Mutex
}

func (c *Core) shutdown() error {
	shutdownTimeout := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	var wg sync.WaitGroup

	done := make(chan struct{})
	for _, shI := range c.onShutdown {
		wg.Add(1)
		go func(shI Shutdowner) {
			defer wg.Done()
			err := shI.Shutdown(ctx)
			if err != nil {
				c.logger.Error(err.Error())
			}
		}(shI)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func (t *Server) Serve() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", t.conf.Host, t.conf.Port),
	)
	if err != nil {
		return err
	}

	return t.Server.Serve(listener)
}

func New(coreApp *core.Core, c Config) *App {
	return &App{
		core:   coreApp,
		config: c,
	}
}

func (a *App) Init() (err error) {
	dbc := NewDB(a.config.Database)
	err = dbc.Init()
	if err != nil {
		return fmt.Errorf("App.InitDB: %w", err)
	}
	
	dbConf := a.config.Database
	conn := fmt.Sprint("postgres://%s:%s@%s:%d/%s", dbConf.user, dbConf.pass, dbConf.host, dbConf.port, dbConf.dbname)
	dbConn, err := sqlx.Connect("postgres", conn)
    if err != nil {
        log.Fatalln(err)
    }

	_, err = conn.Version()
	if err != nil {
		return fmt.Errorf("App.Db: %w", err)
	}

	productRepo := db.NewProductRepo(conn)

	// a.core.RegisterShutdowner(a)

	return
}

func (a *App) Shutdown(context.Context) error {
	a.planner.Stop()

	return nil
}
