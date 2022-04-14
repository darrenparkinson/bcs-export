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

func listDevices(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.logger.Info().Msg("retrieving devices")
	res, err := app.bcs.ListAllDevices(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing devices")
		return err
	}
	printDevices(res)
	return nil
}

func listAssets(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.logger.Info().Msg("retrieving assets")
	res, err := app.bcs.ListAllAssets(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing devices")
		return err
	}
	printAssets(res)
	return nil
}

func printDevices(devices []*ciscobcs.Device) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	fmt.Println()
	tblFields := table.New("ID", "Mgmt IP", "IPv4", "Name", "Status", "Type", "Image", "SW Type", "SW Version", "Product Family", "Product Type")
	tblFields.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, d := range devices {
		tblFields.AddRow(strconv.Itoa(d.GetDeviceId()), d.GetDeviceIp(), d.GetIpAddress(), d.GetDeviceName(), d.GetDeviceStatus(), d.GetDeviceType(), d.GetImageName(), d.GetSwType(), d.GetSwVersion(), d.GetProductFamily(), d.GetProductType())
	}
	tblFields.Print()
	fmt.Println()
}

func printAssets(assets []*ciscobcs.Asset) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	fmt.Println()
	tblFields := table.New("ID", "Device Name", "Chassis Name", "Physical Element ID", "Product Type", "Serial Number")
	tblFields.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, a := range assets {
		tblFields.AddRow(strconv.Itoa(a.GetDeviceId()), a.GetDeviceName(), a.GetChassisName(), a.GetPhysicalElementId(), a.GetProductType(), a.GetSerialNumber())
	}
	tblFields.Print()
	fmt.Println()
}

func uploadDevices(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	devices, err := app.bcs.ListAllDevices(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing devices")
		return err
	}
	app.logger.Info().Int("count", len(devices)).Msg("retrieved devices from cisco bcs")

	uploadObjectsToDatabase(app.db, devices)
	return nil
}

func uploadAssets(c *cli.Context) error {
	app := c.Context.Value("app").(*application)
	app.mustLoadDatabaseConfig()
	defer app.sqlDB.Close()

	assets, err := app.bcs.ListAllAssets(context.Background())
	if err != nil {
		app.logger.Err(err).Msg("received error listing assets")
		return err
	}
	app.logger.Info().Int("count", len(assets)).Msg("retrieved assets from cisco bcs")

	uploadObjectsToDatabase(app.db, assets)
	return nil
}
