package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/darrenparkinson/bcs-export/pkg/ciscobcs"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

func uploadRiskMigitationSummaries(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	summaries, err := app.bcs.ListRiskMitigationSummaries(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing risk mitigation summaries")
		return err
	}
	app.logger.Info().Int("count", len(summaries)).Msg("retrieved risk mitigation summaries from cisco bcs")

	uploadObjectsToDatabase(app.db, summaries)
	return nil
}

func uploadRiskMigitationDetails(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	details, err := app.bcs.ListRiskMitigationDetails(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing risk mitigation details")
		return err
	}
	app.logger.Info().Int("count", len(details)).Msg("retrieved risk mitigation details from cisco bcs")

	uploadObjectsToDatabase(app.db, details)
	return nil
}

func listRiskMitigationSummaries(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.logger.Info().Msg("retrieving risk mitigation summaries")
	res, err := app.bcs.ListRiskMitigationSummaries(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing risk mitigation summaries")
		return err
	}
	printRiskMitigationSummaries(res)
	return nil
}

func listRiskMitigationDetails(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.logger.Info().Msg("retrieving risk mitigation details")
	res, err := app.bcs.ListRiskMitigationDetails(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing risk mitigation details")
		return err
	}
	printRiskMitigationDetails(res)
	return nil
}

func printRiskMitigationSummaries(summaries []*ciscobcs.RiskMitigationSummary) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	fmt.Println()
	tblFields := table.New("High Risk Device Count", "Medium Risk Device Count", "Low Risk Device Count", "Product Family")
	tblFields.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, s := range summaries {
		// tblFields.AddRow(strconv.Itoa(d.GetDeviceId()), d.GetDeviceIp(), d.GetIpAddress(), d.GetDeviceName(), d.GetDeviceStatus(), d.GetDeviceType(), d.GetImageName(), d.GetSwType(), d.GetSwVersion(), d.GetProductFamily(), d.GetProductType())
		tblFields.AddRow(s.GetHighRiskDeviceCount(), s.GetMediumRiskDeviceCount(), s.GetLowRiskDeviceCount(), s.GetProductFamily())
	}
	tblFields.Print()
	fmt.Println()
}

func printRiskMitigationDetails(details []*ciscobcs.RiskMitigationDetails) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	fmt.Println()
	tblFields := table.New("Device ID", "Device IP", "Device Name", "Product Family", "Product ID", "Risk Category", "Risk Score", "SW Type", "SW Version")
	tblFields.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, d := range details {
		tblFields.AddRow(strconv.Itoa(d.GetDeviceId()), d.GetDeviceIp(), d.GetDeviceName(), d.GetProductFamily(), d.GetProductId(), d.GetRiskCategory(), d.GetRiskScore(), d.GetSwType(), d.GetSwVersion())
	}
	tblFields.Print()
	fmt.Println()
}
