package main

import (
	"github.com/Eaacisternas/pokeBackRipley/libs"
)

func main() {
	libs.S3.Region = "us-east-1"
	libs.S3.NewSession(libs.S3.Region)
	libs.S3.Ls()
}
