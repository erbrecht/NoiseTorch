// +build distro

package main

import (
	"fmt"
	"log"
)

func loadLadspaInput(ctx *ntcontext, inp *device) error {

	c := ctx.paClient

	idx, err := c.LoadModule("module-ladspa-sink",
		fmt.Sprintf("sink_name=nui_mic_raw_in sink_master=nui_mic_denoised_out "+
			"label=noisetorch plugin=%s control=%d", "rnnoise_ladspa.so", ctx.config.Threshold))
	if err != nil {
		return err
	}
	log.Printf("Loaded ladspa sink as idx: %d\n", idx)

	return nil
}

func loadLadspaOutput(ctx *ntcontext, out *device) error {

	c := ctx.paClient

	_, err := c.LoadModule("module-ladspa-sink", fmt.Sprintf(`sink_name=nui_out_ladspa sink_master=nui_out_out_sink `+
		`label=noisetorch channels=1 plugin=%s control=%d rate=%d`,
		"rnnoise_ladspa.so", ctx.config.Threshold, 48000))
	if err != nil {
		return err
	}

	return nil
}
