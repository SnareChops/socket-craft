package docker

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Replace(project string) error {
	return cmd.Run("docker-compose", "-f", "local.yml", "up", "-d", "--no-deps", project)
}
