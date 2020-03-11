package docker

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Logs(project string) error {
	return cmd.Run("docker-compose", "-f", "local.yml", "logs", project)
}
