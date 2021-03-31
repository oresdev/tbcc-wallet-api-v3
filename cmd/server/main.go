package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/conf"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/router"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/store"
)

const appName = "tbcc-wallet-api-v3"

func main() {
	c, err := conf.ParseConfig(appName)
	if err != nil {
		logrus.Fatalf("parsing config; %v", err)
	}

	connectStr := fmt.Sprintf(c.DB.Tmpl, c.DB.Host, c.DB.Port, c.DB.Name, c.DB.User, c.DB.Password, c.DB.Schema, appName)

	db, err := store.СreateDB(connectStr, c.DB.ConnLifetime, c.DB.MaxIdleConns, c.DB.PoolSize)
	if err != nil {
		logrus.Errorf("opening connection: %v", err)
	}
	logrus.RegisterExitHandler(func() {
		db.Close()
	})

	handler, err := router.CreateHTTPHandler(db)
	if err != nil {
		logrus.Fatalf("creating http handler: %v", err)
	}

	listenErr := make(chan error, 1)
	server := &http.Server{
		Addr:    ":" + os.Args[1],
		Handler: handler,
	}

	go func() {
		logrus.Println("server started at port:", os.Args[1], time.Now().Format(time.RFC3339))
		listenErr <- server.ListenAndServe()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-listenErr:
		logrus.Fatal(err)
	case <-osSignals:
		server.SetKeepAlivesEnabled(false)
		timeout := time.Second * 5

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
		}
		logrus.Println("stop server")
		logrus.Exit(0)
	}
}
