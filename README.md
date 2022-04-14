# Cisco BCS (Cisco Business Critical Services) Export

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/bcs-export) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/bcs-export) ![GitHub](https://img.shields.io/github/license/darrenparkinson/bcs-export?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/bcs-export)](https://pkg.go.dev/github.com/darrenparkinson/bcs-export) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/bcs-export)](https://goreportcard.com/report/github.com/darrenparkinson/bcs-export)

This repository consists of a library and CLI utility for extracting information from the Cisco BCS Insights API.

* Docs for this API here: https://api.csco-bcs.com/v2/
* Postman Collection here: https://www.getpostman.com/collections/c2abd62e08b6d27a7116 

## Using the CLI

### Installation

If you have Go, you can install the CLI with:

```sh
$ go install github.com/darrenparkinson/bcs-export/cmd/bcs-cli@latest
```

Otherwise you can download a binary for your platform from the [releases](https://github.com/darrenparkinson/bcs-export/releases) page.  

### Configuration

In order to use the CLI, you will need the following environment variables:

* `BCS_APIKEY` - API Key
* `BCS_CUSTOMERID` - Customer ID, e.g. 280987866
* `BCS_DSN` - Database connection details, only required for upload command. e.g. `sqlserver://<user>:<password>@<host>?database=<dbname>`

You may also put them in a `.env` file.

### Usage

You can see detailed help as follows:

* `$ bcs-cli list --help`
* `$ bcs-cli upload --help`

The cli will allow you to list items and upload items to a Microsoft SQL Database (this could be expanded to support additional databases).

```sh
$ bcs-cli -h

NAME:
   bcs-cli - Cisco Business Critical Services (aka BCS) insights data CLI

USAGE:
   bcs-cli [global options] command [command options] [arguments...]

COMMANDS:
   list, l    list various objects
   upload, u  upload objects to the configured database
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d  enable debug logs (default: false)
   --help, -h   show help (default: false)
```

## Using the library

```sh
go get github.com/darrenparkinson/bcs-export/pkg/ciscobcs
```
