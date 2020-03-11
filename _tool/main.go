package main

import (
	"os"

	"github.com/SnareChops/socket-craft/tool/docker"
	"github.com/SnareChops/socket-craft/tool/npx"
)

func main() {
	if len(os.Args) < 2 {
		Help()
	}

	switch os.Args[1] {
	case "up":
		check(docker.Up())
		break
	case "down":
		check(docker.Down())
		break
	case "site":
		check(npx.Run("tsc"))
		check(npx.Run("webpack"))
		check(npx.Run("sass", "sass/index.scss", "_dist/index.css"))
		check(docker.Build("site"))
		check(docker.Replace("site"))
		break
	case "server":
		check(docker.Build("server"))
		check(docker.Replace("server"))
		break
	case "nexus":
		check(docker.Build("nexus"))
		check(docker.Replace("nexus"))
		break
	default:
		Help()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
