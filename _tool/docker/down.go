package docker

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Down() error {
	return cmd.Run("docker-compose", "-f", "local.yml", "down")
}
