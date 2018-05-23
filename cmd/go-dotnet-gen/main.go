package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/matiasinsaurralde/go-dotnet/generator"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	goExtension = ".go"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
	source  = kingpin.Arg("input", "Parse input and generate bindings based on it").ExistingFileOrDir()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	if *source == "" {
		kingpin.Usage()
		return
	}
	var input []string
	fileInfo, _ := os.Stat(*source)
	if fileInfo.IsDir() {
		files, _ := ioutil.ReadDir(*source)
		os.Chdir(*source)
		for _, v := range files {
			fname := v.Name()
			if ext := filepath.Ext(fname); ext != goExtension {
				continue
			}
			input = append(input, fname)
		}
	} else {
		input = []string{*source}
	}
	g := generator.New(input).Verbose(*verbose)
	err := g.Parse()
	if err != nil {
		panic(err)
	}
	g.Generate()
}
