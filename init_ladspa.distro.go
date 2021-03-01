// +build distro

package main

func init_ladspa(ctx *ntcontext) {
	// noop for distro builds
	// rnnoise_ladspa.so should be installed with the package,
	// no need to load embedded version
}
