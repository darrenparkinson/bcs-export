package main

import (
	"context"

	"github.com/urfave/cli/v2"
)

func uploadProductAlerts(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()
	ctx := context.Background()
	switch c.Command.Name {

	case "field-notices":
		app.logger.Info().Msg("retrieving field-notices from cisco bcs")
		items, err := app.bcs.ListAllFieldNotices(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "field-notice-bulletins":
		app.logger.Info().Msg("retrieving field-notice-bulletins from cisco bcs")
		items, err := app.bcs.ListAllFieldNoticeNBulletins(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "hweox":
		app.logger.Info().Msg("retrieving hweox from cisco bcs")
		items, err := app.bcs.ListAllHardwareEOX(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "hweox-bulletins":
		app.logger.Info().Msg("retrieving hweox-bulletins from cisco bcs")
		items, err := app.bcs.ListAllHardwareEOXBulletins(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "sweox":
		app.logger.Info().Msg("retrieving sweox from cisco bcs")
		items, err := app.bcs.ListAllSoftwareEOX(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "sweox-bulletins":
		app.logger.Info().Msg("retrieving sweox-bulletins from cisco bcs")
		items, err := app.bcs.ListAllSoftwareEOXBulletins(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "psirts":
		app.logger.Info().Msg("retrieving psirts from cisco bcs")
		items, err := app.bcs.ListAllSecurityAdvisories(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "psirt-bulletins":
		app.logger.Info().Msg("retrieving psirt-bulletins from cisco bcs")
		items, err := app.bcs.ListAllSecurityAdvisoryBulletins(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	case "software-alerts":
		app.logger.Info().Msg("retrieving software-alerts from cisco bcs")
		items, err := app.bcs.ListAllSoftwareAlerts(ctx)
		if err != nil {
			app.logger.Err(err).Msg("received error retrieving items")
			return err
		}
		app.logger.Info().Int("count", len(items)).Msg("retrieved items from cisco bcs")
		uploadObjectsToDatabase(app.db, items)

	default:
		app.logger.Error().Msg("invalid command")
	}

	return nil
}
