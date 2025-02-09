package vtest

import (
	plugin "github.com/vg006/vtest/internal/plugins"
)

type Options struct {
	Path    string
	DevMode bool
	Exclude []string
}

func Expose(router interface{}) {
	switch TypeOf(router) {
	case "gin":
		plugin.Gin()
	case "fiber":
		plugin.Fiber()
	case "chi":
		plugin.Chi()
	case "echo":
		plugin.Echo()
	default:
		return
	}
}
