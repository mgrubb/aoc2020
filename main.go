package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/mgrubb/aoc2020/plugin"
)

var (
	printPlugins bool
	sampleOnly   bool
)

func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func inputIsSample(input string) bool {
	return strings.HasSuffix(input, "sample.txt")
}

func findInputs(pluginName string) ([]string, error) {
	files, err := os.ReadDir(path.Join("inputs", pluginName))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("No input found for %s", pluginName)
	}
	inputs := make([]string, len(files))
	for i, v := range files {
		if v.Type().IsRegular() {
			inputs[i] = path.Join("inputs", pluginName, files[i].Name())
		}
	}

	// Sort sample.txt first
	sort.SliceStable(inputs, func(i, j int) bool {
		if inputIsSample(inputs[i]) {
			return true
		}
		return inputs[i] < inputs[j]
	})

	return inputs, nil
}

func runPlugin(p plugin.Plugin) error {
	inputs, err := findInputs(p.Name())
	must(err)
	if sampleOnly {
		if !inputIsSample(inputs[0]) {
			return fmt.Errorf("Sample only run requested, but no sample found for %s", p.ProperName())
		}
		inputs = inputs[:1]
	}
	for _, input := range inputs {
		inr, err := os.Open(input)
		defer inr.Close()
		must(err)
		fmt.Printf("%s input: %s\n", p.ProperName(), input)
		err = p.Solve(inr)
		must(err)
	}
	return nil
}

func runPlugins(plugins []plugin.Plugin) {
	var err error
	for _, p := range plugins {
		if printPlugins {
			fmt.Printf("%s: %s\n", p.Name(), p.ProperName())
		} else {
			err = runPlugin(p)
			must(err)
		}
	}
}

func parsePluginName(arg string) string {
	if strings.HasPrefix(arg, "day") {
		return arg
	}
	i, err := strconv.Atoi(arg)
	must(err)
	return fmt.Sprintf("day%02d", i)
}

func main() {
	must(plugin.PluginManager.LoadPlugins("plugins"))

	flag.BoolVar(&printPlugins, "plugins", false, "Show list of plugins")
	flag.BoolVar(&sampleOnly, "sample", false, "Only run sample input")
	flag.Parse()
	nonFlags := flag.Args()

	var plugins []plugin.Plugin
	if len(nonFlags) > 0 {
		for _, arg := range nonFlags {
			pluginName := parsePluginName(arg)
			p, ok := plugin.PluginManager.PluginByName(pluginName)
			if !ok {
				must(fmt.Errorf("Not plugin named %s found", arg))
			}
			plugins = append(plugins, p)
		}
	} else {
		plugins = plugin.PluginManager.Plugins()
	}

	runPlugins(plugins)

	os.Exit(0)
}
