package docker

import "github.com/SnareChops/socket-craft/_tool/cmd"

func Build(project string) error {
	return cmd.Run("docker", "build", "--target", "release", "-t", "socket-craft/"+project, "-f", project+".docker", ".")
}
