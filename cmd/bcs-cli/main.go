package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/darrenparkinson/bcs-export/internal/helpers"
	"github.com/darrenparkinson/bcs-export/pkg/ciscobcs"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type application struct {
	logger     *zerolog.Logger
	bcs        *ciscobcs.Client
	customerID string
	dsn        string
	db         *gorm.DB
	sqlDB      *sql.DB
}

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.SetFlags(0)
	log.SetOutput(logger)
	app := &cli.App{
		Name:  "bcs-cli",
		Usage: "Cisco Business Critical Services (aka BCS) insights data CLI",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "enable debug logs",
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			} else {
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			}
			// Set as console writer always since it's a cli
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
			godotenv.Load() // enable loading from .env if available
			var customerID, apikey, dsn string
			helpers.MustMapEnv(&customerID, "BCS_CUSTOMERID")
			helpers.MustMapEnv(&apikey, "BCS_APIKEY")
			dsn = os.Getenv("BCS_DSN")
			bcs, err := ciscobcs.NewClient(customerID, apikey, nil)
			if err != nil {
				return fmt.Errorf("error initialising bcs client: %w", err)
			}
			c.Context = context.WithValue(c.Context, "app", &application{
				logger:     &logger,
				bcs:        bcs,
				customerID: customerID,
				dsn:        dsn,
			})
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list various objects",
				Subcommands: []*cli.Command{
					{
						Name:    "devices",
						Aliases: []string{"d"},
						Usage:   "list devices",
						Action:  listDevices,
					},
					{
						Name:    "assets",
						Aliases: []string{"a"},
						Usage:   "list assets",
						Action:  listAssets,
					},
				},
			},
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "upload objects to the configured database",
				Subcommands: []*cli.Command{
					{
						Name:    "cbp",
						Aliases: []string{"c"},
						Usage:   "upload configuration best practices",
						Action:  uploadCBP,
					},
					{
						Name:    "inventory",
						Aliases: []string{"i"},
						Usage:   "upload inventory items",
						Subcommands: []*cli.Command{
							{
								Name:    "devices",
								Aliases: []string{"d"},
								Usage:   "upload devices",
								Action:  uploadDevices,
							},
							{
								Name:    "assets",
								Aliases: []string{"a"},
								Usage:   "upload assets",
								Action:  uploadAssets,
							},
						},
					},
					{
						Name:    "product-alerts",
						Aliases: []string{"p"},
						Usage:   "upload product alert items individually",
						Subcommands: []*cli.Command{
							{
								Name:    "field-notices",
								Aliases: []string{"fn"},
								Usage:   "upload field notices",
								Action:  uploadProductAlerts,
							},
							{
								Name:    "field-notice-bulletins",
								Aliases: []string{"fnbulletins"},
								Usage:   "upload field notice bulletins",
								Action:  uploadProductAlerts,
							},
							{
								Name:   "hweox",
								Usage:  "upload harware eox",
								Action: uploadProductAlerts,
							},
							{
								Name:    "hweox-bulletins",
								Aliases: []string{"hweoxbulletins"},
								Usage:   "upload hardware eox bulletins",
								Action:  uploadProductAlerts,
							},
							{
								Name:    "psirts",
								Aliases: []string{"security-advisories"},
								Usage:   "upload security advisories",
								Action:  uploadProductAlerts,
							},
							{
								Name:    "psirt-bulletins",
								Aliases: []string{"security-advisory-bulletins"},
								Usage:   "upload psirt bulletins",
								Action:  uploadProductAlerts,
							},
							{
								Name:    "software-alerts",
								Aliases: []string{"sw"},
								Usage:   "upload software alerts",
								Action:  uploadProductAlerts,
							},
							{
								Name:   "sweox",
								Usage:  "upload software eox",
								Action: uploadProductAlerts,
							},
							{
								Name:    "sweox-bulletins",
								Aliases: []string{"sweoxbulletins"},
								Usage:   "upload software eox bulletins",
								Action:  uploadProductAlerts,
							},
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal().Err(err).Msg("cli error")
	}
}
