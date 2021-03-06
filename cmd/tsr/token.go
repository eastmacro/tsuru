// Copyright 2014 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/tsuru/tsuru/app"
	"github.com/tsuru/tsuru/cmd"
)

type tokenCmd struct{}

func (tokenCmd) Run(context *cmd.Context, client *cmd.Client) error {
	t, err := app.AuthScheme.AppLogin("tsr")
	if err != nil {
		return err
	}
	fmt.Fprintf(context.Stdout, t.GetValue())
	return nil
}

func (tokenCmd) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "token",
		Usage:   "token",
		Desc:    "Generates a tsuru token.",
		MinArgs: 0,
	}
}
