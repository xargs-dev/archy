package source

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"go.xargs.dev/archy"
	"go.xargs.dev/archy/option"
	"go.xargs.dev/archy/uname"
)

// We _can_ make this configurable but likely nobody needs it?
const unamePath = "uname"

type Uname struct{}

func (u *Uname) Values(ctx context.Context, o *option.Option, v *archy.Values) error {
	v.Source = unamePath

	if o.KernelName {
		val, err := u.cmd(ctx, uname.FlagKernelName)
		if err != nil {
			return fmt.Errorf("reading kernel name from uname: %w", err)
		}
		v.KernelName = val
	}
	if o.Machine {
		val, err := u.cmd(ctx, uname.FlagMachine)
		if err != nil {
			return fmt.Errorf("reading machine from uname: %w", err)
		}
		v.Machine = val
	}
	return nil
}

func (*Uname) cmd(ctx context.Context, flag string) (string, error) {
	cmd := exec.CommandContext(ctx, unamePath, "--"+flag)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(out)
	return strings.TrimSpace(str), nil
}
