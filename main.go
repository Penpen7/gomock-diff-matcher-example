package main

import (
	"github.com/Penpen7/gomock-diff-matcher-example/cmd"
	"github.com/Penpen7/gomock-diff-matcher-example/repository"
)

func main() {

	cmd.Run(repository.NewOnMemoryUserRepository(), 1)
}
