package main

import (
	"fmt"
	"os"

	"github.com/digiconvent/d9t/meta"
	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/flags"
	"github.com/digiconvent/install_on_debian"
)

func main() {
	meta.Initialise(os.Args)
	bin := install_on_debian.NewBinary("digiconvent")

	switch *flags.Mode {
	case "install":
		if bin.IsInstalled() {
			panic("cannot install when already installed")
		}
		environment.Env.Prompt()
		// save in the binary
		err := environment.Save()
		if err != nil {
			panic(err)
		}

		_, err = bin.Install()
		if err != nil {
			panic(err)
		}
	case "serve":
		fmt.Println("serve")
	}
}
