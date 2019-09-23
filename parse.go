package dockerw

import (
	"encoding/json"
	"errors"
)

type Params struct {
	Cmd   string // noop, metadata, version, start
	Pairs []string
}

var (
	metadataArg = "docker-cli-plugin-metadata"
	argsErr     = errors.New("invalid args")
)

// Metadata returns the metadata string
func Metadata() string {
	data := struct {
		SchemaVersion, Vendor, Version, ShortDescription string
	}{
		"0.1.0",
		"https://github.com/karlpokus/docker-watch",
		Version,
		"watch and reload services",
	}
	b, _ := json.Marshal(data)
	return string(b)
}

// isMetadata returns whether input contains metadataArg
func isMetadata(args []string) bool {
	for _, v := range args {
		if v == metadataArg {
			return true
		}
	}
	return false
}

// Args parses cli args into a Params
func Args(args []string) (p Params, err error) {
	if isMetadata(args) {
		p.Cmd = "metadata"
		return
	}
	n := len(args)
	if n == 2 {
		p.Cmd = "noop"
		return
	}
	if n == 3 && args[2] == "version" {
		p.Cmd = args[2]
		return
	}
	if n >= 4 { // start pairs...
		p.Cmd = args[2]
		p.Pairs = args[3:]
		return
	}
	err = argsErr
	return
}
