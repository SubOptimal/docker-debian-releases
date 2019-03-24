//usr/bin/env go run "$0" "$@"; exit "$?"
// -*- go -*-

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

func panicIf(err error, transforms ...func(error) error) {
	if err == nil {
		return
	}
	for _, transform := range transforms {
		err = transform(err)
	}
	panic(err)
}

func cmdLineParse() (string, string, string, error) {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		return "", "", "", fmt.Errorf("could not parse %s", os.Args)
	}
	output := "-"
	if len(os.Args) == 4 {
		output = os.Args[3]
	}
	return os.Args[1], os.Args[2], output, nil
}

type releaseInfo struct {
	Origin        string
	Label         string
	Suite         string
	Version       string
	Date          time.Time
	Codename      string
	URL           string
	Architectures []string
}

func main() {
	jsonfilename, tmplfilename, outfilename, err := cmdLineParse()
	panicIf(err)
	jsonString, err := ioutil.ReadFile(jsonfilename)
	panicIf(err)
	var releaseInfos []releaseInfo
	err = json.Unmarshal(jsonString, &releaseInfos)
	panicIf(err)
	t, err := template.New("").Funcs(template.FuncMap{
		"stringsIndex": strings.Index,
	}).ParseFiles(tmplfilename)
	panicIf(err)
	if outfilename != "-" {
		outfile, err := os.Create(outfilename)
		panicIf(err)
		err = t.ExecuteTemplate(outfile, tmplfilename, releaseInfos)
		panicIf(err)
	} else {
		err = t.ExecuteTemplate(os.Stdout, tmplfilename, releaseInfos)
		panicIf(err)
	}
}
