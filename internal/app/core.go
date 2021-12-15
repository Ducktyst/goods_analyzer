package app

type Core struct {
	config  Config
	version string
	service string
}

func NewCore(version, service string, cfg interface{}) *Core {
	c := &Core{
		version: version,
		service: service,
	}
	c.init(cfg)

	return c
}

func (c *Core) Version() string {
	return c.version
}

func (c *Core) Service() string {
	return c.service
}

func (c *Core) init(cfg interface{}) *Core {

	//if configValidator, ok := cfg.(ConfigValidator); ok {
	//	if err := configValidator.Validate(); err != nil {
	//		c.logger.Fatal(err)
	//	}
	//}
	//
	//if configGetter, ok := cfg.(ConfigGetter); ok {
	//	c.config = configGetter.GetConfig()
	//
	//	if c.config.Sentry.DSN != "" {
	//		if err = sentry.Init(sentry.ClientOptions{
	//			Dsn:         c.config.Sentry.DSN,
	//			Environment: c.config.Sentry.Environment,
	//			Release:     c.version,
	//		}); err != nil {
	//			c.logger.Fatal(err)
	//		}
	//	}
	//}
	//
	//c.tracer = &TracingService{
	//	conf: c.config.Tracing,
	//}
	//c.RegisterInitializer(c.tracer)
	//c.RegisterShutdowner(c.tracer)
	//
	//c.metrics = &MetricsService{
	//	conf: c.config.Metrics,
	//}
	//c.RegisterInitializer(c.metrics)
	//
	//c.adminServer.Use(sentryecho.New(sentryecho.Options{}))

	return c
}

func (c *Core) Run(app App) {
	//c.RegisterInitializer(app)
	//for _, i := range c.initializers {
	//	if err := i.Init(); err != nil {
	//		c.logger.Fatal(err)
	//	}
	//}
	//
	//c.adminServer.GET("/health", c.healthCheckHandler)
	//c.adminServer.Any("/metrics", echo.WrapHandler(c.metrics))
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	//
	//// run app and send panic to sentry
	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			sentry.CurrentHub().Recover(err)
	//			sentry.Flush(time.Second * 3)
	//			c.logger.Panic(err)
	//		}
	//	}()
	//
	//	if err := app.Run(); err != nil {
	//		sentry.CaptureException(err)
	//		c.logger.Error(err)
	//		quit <- syscall.SIGTERM
	//	}
	//}()
	//
	//go func() {
	//	err := c.adminServer.Start(fmt.Sprintf("%s:%d", c.config.AdminServer.Host, c.config.AdminServer.Port))
	//	if err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		c.logger.Panic("http server err=%q", err)
	//	}
	//}()
	//
	//<-quit
	//if err := c.shutdown(); err != nil {
	//	c.logger.Fatal(err)
	//}
}
