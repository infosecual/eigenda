package main

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_isLocalhost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, socket string) {
		isLocalhost(socket)
	})
}

func Fuzz_Nosy_pluginOps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		pluginOps(ctx)
	})
}
