package main

import (
	"github.com/sethvargo/go-githubactions"
)

const cloudDomainPublic = "https://gh-api.atlasgo.cloud"

func main() {
	act := githubactions.New()
	act.Infof("Hello, world!")
}
