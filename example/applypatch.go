package main

import (
	"os"

	"github.com/mrshubhambasu/xupdate/integrator"
)

func main() {
	if len(os.Args) != 4 {
		printUsage(1)
	}

	//initialize new update integrator
	i := integrator.New()

	err := i.Integrate(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		printUsage(1)
		panic(err)
	}
}

func printUsage(exitcode int) {
	println("usage: " + os.Args[0] + " oldfile newfile patchfile")
	os.Exit(exitcode)
}
