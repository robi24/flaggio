package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	ApplicationName        = "flaggio"
	ApplicationDescription = "Self hosted feature flag solution"
	ApplicationVersion     = "0.1.0"
	GitSummary             = ""
	GitBranch              = ""
	BuildStamp             = ""
)

const (
	logFormatterText = "text"
	logFormatterJSON = "json"
)

func build() string {
	return fmt.Sprintf("%s[%s] (%s)", GitBranch, GitSummary, BuildStamp)
}

func main() {
	app := cli.App{
		Name:        ApplicationName,
		Description: ApplicationDescription,
		Version:     ApplicationVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database-uri",
				Usage:   "Database URI",
				EnvVars: []string{"DATABASE_URI"},
				Value:   "mongodb://localhost:27017/flaggio",
			},
			&cli.StringFlag{
				Name:    "redis-uri",
				Usage:   "Redis URI",
				EnvVars: []string{"REDIS_URI"},
				Value:   "redis://localhost:6379",
			},
			&cli.StringFlag{
				Name:    "build-path",
				Usage:   "UI build absolute path",
				EnvVars: []string{"BUILD_PATH"},
				Value:   "/",
			},
			&cli.StringSliceFlag{
				Name:    "cors-allowed-origins",
				Usage:   "CORS allowed origins separated by comma",
				EnvVars: []string{"CORS_ALLOWED_ORIGINS"},
			},
			&cli.StringSliceFlag{
				Name:    "cors-allowed-headers",
				Usage:   "CORS allowed headers",
				EnvVars: []string{"CORS_ALLOWED_HEADERS"},
			},
			&cli.BoolFlag{
				Name:    "cors-debug",
				Usage:   "CORS debug logging",
				EnvVars: []string{"CORS_DEBUG"},
				Hidden:  true,
			},
			&cli.BoolFlag{
				Name:    "no-api",
				Usage:   "Don't start the API server",
				EnvVars: []string{"NO_API"},
			},
			&cli.BoolFlag{
				Name:    "no-admin",
				Usage:   "Don't start the admin server",
				EnvVars: []string{"NO_ADMIN"},
			},
			&cli.BoolFlag{
				Name:    "no-admin-ui",
				Usage:   "Don't start the admin UI",
				EnvVars: []string{"NO_ADMIN_UI"},
			},
			&cli.StringFlag{
				Name:    "app-env",
				Usage:   "Environment this application is running on. Valid values are: dev, production",
				EnvVars: []string{"APP_ENV"},
				Value:   "production",
			},
			&cli.StringFlag{
				Name:    "api-addr",
				Usage:   "Sets the bind address for the API",
				EnvVars: []string{"API_ADDR"},
				Value:   ":8080",
			},
			&cli.StringFlag{
				Name:    "admin-addr",
				Usage:   "Sets the bind address for the admin",
				EnvVars: []string{"ADMIN_ADDR"},
				Value:   ":8081",
			},
			&cli.StringFlag{
				Name:    "log-formatter",
				Usage:   "Sets the log formatter for the application. Valid values are: text, json",
				EnvVars: []string{"LOG_FORMATTER"},
				Value:   logFormatterJSON,
			},
			&cli.StringFlag{
				Name:    "log-level",
				Usage:   "Sets the log level for the application",
				EnvVars: []string{"LOG_LEVEL"},
				Value:   "info",
			},
		},
		Action: func(c *cli.Context) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			logger := logrus.New()
			logLevel, err := logrus.ParseLevel(c.String("log-level"))
			if err != nil {
				return err
			}
			logger.SetLevel(logLevel)
			switch c.String("log-formatter") {
			case logFormatterText:
				logger.SetFormatter(new(logrus.TextFormatter))
			case logFormatterJSON:
				logger.SetFormatter(new(logrus.JSONFormatter))
			default:
				return fmt.Errorf("invalid formatter: %s", c.String("log-formatter"))
			}

			logger.
				WithFields(logrus.Fields{"version": ApplicationVersion, "build": build()}).
				Infof("starting %s application", ApplicationName)

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			errs := make(chan error, 1)
			if !c.Bool("no-api") {
				go func() {
					err := startAPI(ctx, c, logger.WithField("app", "api"))
					if err != nil {
						errs <- err
					}
				}()
			}
			if !c.Bool("no-admin") {
				go func() {
					err := startAdmin(ctx, c, logger.WithField("app", "admin"))
					if err != nil {
						errs <- err
					}
				}()
			}

			for {
				select {
				case err := <-errs:
					return err
				case <-done:
					cancel()
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
