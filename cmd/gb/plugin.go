package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/constabulary/gb/cmd"
)

func init() {
	registerCommand(PluginCmd)
}

var PluginCmd = &cmd.Command{
	Name:  "plugin",
	Short: "plugin information",
	Long: `gb supports git style plugins.

A gb plugin is anything in the $PATH with the prefix gb-. In other words
gb-something, becomes gb something.

gb plugins are executed from the parent gb process with the environment
variable, GB_PROJECT_DIR set to the root of the current project.

gb plugins can be executed directly but this is rarely useful, so authors
should attempt to diagnose this by looking for the presence of the 
GB_PROJECT_DIR environment key.
`,
}

func lookupPlugin(arg string) (string, error) {
	plugin := "gb-" + arg
	path, err := exec.LookPath(plugin)
	if err != nil {
		return "", fmt.Errorf("plugin: unable to locate %q: %v", plugin, err)
	}
	return path, nil
}

func findPlugins(paths []string) ([]string, error) {
	pluginNames := []string{}

	for _, path := range paths {
		filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			bName := filepath.Base(path)
			if len(bName) > 3 && bName[0:3] == "gb-" {
				pluginNames = append(pluginNames, bName[3:len(bName)])
			}
			return nil
		})
	}
	return pluginNames, nil
}
