package npx

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Run(args ...string) error {
	args = append([]string{"npx"}, args...)
	return cmd.Run(args...)
}
