package main

import (
	"os"

	"github.com/SnareChops/socket-craft/_tool/docker"
	"github.com/SnareChops/socket-craft/_tool/npx"
)

func main() {
	if len(os.Args) < 2 {
		Help()
	}

	switch os.Args[1] {
	case "init":
		build("nexus")
		build("server")
		build("site")
		up()
		break
	case "up":
		up()
		break
	case "down":
		down()
		break
	case "help":
		Help()
		break
	case "logs":
		logs(os.Args[2])
		break
	default:
		build(os.Args[1])
		replace(os.Args[1])
		break
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func build(project string) {
	switch project {
	case "site":
		check(npx.Run("tsc"))
		check(npx.Run("webpack"))
		check(npx.Run("sass", "site/sass/index.scss", "_dist/index.css"))
		check(docker.Build("site"))
		break
	default:
		check(docker.Build(project))
		break
	}
}

func replace(project string) {
	check(docker.Replace(project))
}

func up() {
	check(docker.Up())
}

func down() {
	check(docker.Down())
}

func logs(project string) {
	check(docker.Logs(project))
}
