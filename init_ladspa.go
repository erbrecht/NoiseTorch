// +build !distro

package main

import (
	"io/ioutil"
)

//go:embed c/ladspa/rnnoise_ladspa.so
var libRNNoise []byte

const appName = "NoiseTorch"

func init_ladspa(ctx *ntcontext) {

	rnnoisefile := dumpLib()
	defer removeLib(rnnoisefile)

	//ctx := ntcontext{}
	ctx.librnnoise = rnnoisefile
}

func dumpLib() string {
	f, err := ioutil.TempFile("", "librnnoise-*.so")
	if err != nil {
		log.Fatalf("Couldn't open temp file for librnnoise\n")
	}
	f.Write(libRNNoise)
	log.Printf("Wrote temp librnnoise to: %s\n", f.Name())
	return f.Name()
}

func removeLib(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Printf("Couldn't delete temp librnnoise: %v\n", err)
	}
	log.Printf("Deleted temp librnnoise: %s\n", file)
}
