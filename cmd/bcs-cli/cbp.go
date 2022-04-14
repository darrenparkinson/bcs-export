package main

import (
	"context"

	"github.com/urfave/cli/v2"
)

func uploadCBP(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	cbpDetails, err := app.bcs.ListAllCBPDetailsPerDeviceID(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing cbp details")
		return err
	}
	app.logger.Info().Int("count", len(cbpDetails)).Msg("retrieved cbp details from cisco bcs")

	uploadObjectsToDatabase(app.db, cbpDetails)

	cbpRules, err := app.bcs.ListAllCBPRules(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing cbp rules")
		return err
	}
	app.logger.Info().Int("count", len(cbpRules)).Msg("retrieved cbp rules from cisco bcs")

	uploadObjectsToDatabase(app.db, cbpRules)

	return nil
}
