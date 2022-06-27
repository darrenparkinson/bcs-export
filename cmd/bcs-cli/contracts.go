package main

import (
	"context"

	"github.com/urfave/cli/v2"
)

func uploadContractSerials(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	serials, err := app.bcs.ListAllContractInformationPerSerialNumber(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing serial contracts")
		return err
	}
	app.logger.Info().Int("count", len(serials)).Msg("retrieved serial contracts from cisco bcs")

	uploadObjectsToDatabase(app.db, serials)

	return nil
}
