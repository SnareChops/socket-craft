package docker

import "github.com/SnareChops/socket-craft/tool/cmd"

func Up() error {
	return cmd.Run("docker-compose", "-f", "local.yml", "up", "-d")
}
