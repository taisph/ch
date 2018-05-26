# ch

ch is a command line tool that makes working with Clubhouse a little more console friendly.

It also contains a self-contained Clubhouse API client that can be integrated by other tools.

Please note that this is still a highly experimental package and all the bits and pieces may be altered in completely
backward-incompatible ways. 

# Installation

## Source

Requirements:
- [Go 1.10 or newer](http://golang.org/doc/install) (probably works with earlier versions but it is currently untested)

To retrieve, build and install the command, run:

```bash
go get -u github.com/taisph/ch/cmd/ch
```

This will place the `ch` command in your `$GOPATH/bin` directory.

# Usage

You need a personal Clubhouse API token for this tool to work its magic. You can generate one
[here](https://app.clubhouse.io/settings/account/api-tokens).

Note that the token is per-organization meaning that if
you work with multiple organizations in Clubhouse, you'll need a token for each one of them.

When you have generated a token you can either pass it using the `--token` parameter or setting the environment variable
`CLUBHOUSE_API_TOKEN`.

## Help

You can get help with the various commands by running any of the following commands:

```bash
ch help
# or
ch story help
# and perhaps
ch story list --help
# alternatively
ch story help list
```

## Listing stories

To list all stories in a given project, run:
```bash
ch story list --project-id 42
```

## Creating stories

To create a new story, run:
```bash
ch story create --project-id 42 "Figure out the answer to everything"
```

You can also create a new story and open it in your browser:
```bash
ch story create --project-id 42 --open "Figure out the answer to everything in a browser"
```
