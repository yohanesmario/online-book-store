package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mikhail-bigun/fiberlogrus"
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"github.com/yohanesmario/online-book-store/conf"
	"github.com/yohanesmario/online-book-store/docs"
	"github.com/yohanesmario/online-book-store/internal/api"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/gormdb"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository"
	"github.com/yohanesmario/online-book-store/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// @title         Online Book Store API
// @version       0.0.1
// @description   API Documentation for Online Book Store.
//
// @contact.name  Yohanes Mario Chandra
// @contact.email yohanes.mc@gmail.com
//
// @externalDocs.description  Github Repository
// @externalDocs.url          https://github.com/yohanesmario/online-book-store
//
// @host      localhost:8080
// @BasePath  /
//
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.<br><b>Example:</b> <code>Bearer &lt;jwt_token_here&gt;</code>
func main() {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", conf.Configuration.APIPort)

	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			logrus.SetOutput(os.Stdout)
			logrus.SetLevel(conf.Configuration.LogLevel)
			logrus.SetFormatter(conf.Configuration.LogFormatter)
			return &fxlogrus.LogrusLogger{Logger: logrus.StandardLogger()}
		}),
		appModule,
		gormdb.GormDBModule,
		repository.RepositoryImplModule,
		service.ServiceModule,

		api.APIModule,
	).Run()
}

var appModule = fx.Module("app",
	fx.Provide(initializeApp),
)

func initializeApp(lifecycle fx.Lifecycle) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
	app.Use(fiberlogrus.New(fiberlogrus.Config{
		Logger: logrus.StandardLogger(),
		Tags: []string{
			fiberlogrus.TagStatus,
			fiberlogrus.TagLatency,
			fiberlogrus.TagMethod,
			fiberlogrus.TagPath,
		},
	}))
	app.Hooks().OnListen(func(data fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}
		scheme := "http"
		if data.TLS {
			scheme = "https"
		}
		logrus.Infof("[APP] Listening on: %s://%s:%s", scheme, data.Host, data.Port)
		return nil
	})

	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			logrus.Info("Starting...")
			go func() {
				err := app.Listen(":" + strconv.Itoa(conf.Configuration.APIPort))
				if err != nil {
					logrus.Errorf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(context context.Context) error {
			logrus.Info("Shutting down...")
			err := app.Shutdown()
			if err != nil {
				logrus.Errorf("Failed to shutdown server: %v", err)
			}
			return err
		},
	})

	return app
}
