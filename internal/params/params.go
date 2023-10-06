package params

import (
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"
)

// ErrHelp is returned when --help flag is
// used and application should not launch.
var ErrHelp = errors.New("help")

// New reads flags and envs and sets it's into Params
// that corresponds to the values read.
func New(params interface{}) error {
	if _, err := flags.Parse(params); err != nil {
		var flagsErr *flags.Error
		if errors.As(err, &flagsErr) && flagsErr.Type == flags.ErrHelp {
			return ErrHelp
		}
		return fmt.Errorf("failed to parse params: %w", err)
	}
	return nil
}

type Params struct {
	Secrets    string `short:"s" long:"secrets" description:"Path to secrets json"`
	DeployPath string `short:"d" long:"deploy" description:"Path to deployment file"`
	OutputPath string `short:"o" long:"output" default:"./.env" description:"Path to output env file"`
}
