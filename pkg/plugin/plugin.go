package plugin

import (
	"path/filepath"
	goplugin "plugin"
)

func LoadPlugins(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	pattern := filepath.Join(abs, "*.so")
	plugins, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, plugin := range plugins {
		if _, err := goplugin.Open(plugin); err != nil {
			return err
		}
	}
	return nil
}
