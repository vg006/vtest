package vtest

import (
	plugin "github.com/vg006/vtest/internal/plugins"
	types "github.com/vg006/vtest/internal/types"
)

func Expose(router interface{}, opts types.Options) {
	if opts.Disable == true {
		return
	}
	if opts.Path == "" {
		opts.Path = "/vtest/routes"
	}

	switch r := router.(type) {
	case types.Fiber:
		plugin.Fiber(r, opts)
	case types.Chi:
		plugin.Chi(r, opts)
	case types.Echo:
		plugin.Echo(r, opts)
	default:
		return
	}
}
