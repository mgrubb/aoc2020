package plugin

import (
	"io"
	"os"
	"path"
	"plugin"
	"sort"
	"strings"
)

// Plugin interface defines what actions a plugin is expected to implement
type Plugin interface {
	Name() string
	ProperName() string
	Solve(input io.Reader) error
}

// Manager struct embodies the plugin manager
type Manager interface {
	LoadPlugins(libpath string) error
	RegisterSolution(plugin Plugin)
	Plugins() []Plugin
	PluginByName(name string) (Plugin, bool)
}

type manager struct {
	plugins []Plugin
}

// PluginManager is the default PluginManager
var PluginManager *manager

func init() {
	PluginManager = &manager{}
}

// LoadPlugins loads the plugins from the filesystem
func (mgr *manager) LoadPlugins(libpath string) error {
	items, err := os.ReadDir(libpath)
	if err != nil {
		return err
	}

	for _, item := range items {
		if !item.Type().IsRegular() {
			continue
		}
		itemName := item.Name()
		if strings.HasSuffix(itemName, ".so") {
			err := mgr.loadPlugin(libpath, itemName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// RegisterSolution adds the plugin into the solution registry
func (mgr *manager) RegisterSolution(plugin Plugin) {
	mgr.plugins = append(mgr.plugins, plugin)
	sort.SliceStable(mgr.plugins, func(i, j int) bool {
		return mgr.plugins[i].Name() < mgr.plugins[j].Name()
	})
}

func (mgr *manager) loadPlugin(libpath, lib string) error {
	pPath := path.Join(libpath, lib)
	p, err := plugin.Open(pPath)
	if err != nil {
		return err
	}
	regFunc, err := p.Lookup("InitSolution")
	if err != nil {
		return err
	}
	regFunc.(func(m Manager))(mgr)
	return nil
}

func (mgr *manager) Plugins() []Plugin {
	return mgr.plugins
}

func (mgr *manager) PluginByName(name string) (Plugin, bool) {
	for _, p := range mgr.plugins {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}
