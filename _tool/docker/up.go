package docker

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Up() error {
	return cmd.Run("docker-compose", "-f", "local.yml", "up", "-d")
}
