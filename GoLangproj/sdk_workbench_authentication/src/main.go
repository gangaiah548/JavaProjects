package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	routeV1 "sdk_workbench_authentication/src/api/routes/v1"
	//routeSDK1 "sdk_workbench_authentication/src/api/routes/SDK1"
	"sdk_workbench_authentication/src/config"
	"sdk_workbench_authentication/src/config/logger"
	"sdk_workbench_authentication/src/errors"
)

func main() {

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := config.Bootstrap()
	props := app.Props

	timeout := time.Duration(props.ResponseTimeout) * time.Second

	// instantiate gin
	gin := config.InitGin(props.AppEnv)

	// setup routes
	routerV1 := gin.Group("v1")
	routeV1.Setup(timeout, routerV1)

	logger.Info().Msg("[👌] Application bootstrap successful!")
	logger.Info().Msg("[🛠️] Attempting to start application on " + props.BindIp + ":" + props.Port)

	// setup server
	srv := &http.Server{
		Addr:    props.BindIp + ":" + props.Port,
		Handler: gin,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(errors.New(err, "[🛑] failure during server startup"))
		}
	}()

	logger.Info().Msg("[✅] Application started on " + props.BindIp + ":" + props.Port)
	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logger.Info().Msg("[⚠️] Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(errors.New(err, "[🛑] Server forced to shutdown! "))
	}

	logger.Info().Msg("[🚩] Server exiting")

	//TODO Error and Exceptions handling
	// ! There is absolutely no Error and exceptions handling here.
	// ! Whoever is looking at this comment, if you have time, please implement
}
