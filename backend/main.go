package main

import (
	"fmt"
	"os"

	api_engine "github.com/digiconvent/d9t/api/engine"
	"github.com/digiconvent/d9t/meta"
	"github.com/digiconvent/d9t/meta/acme"
	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/d9t/meta/flags"
	"github.com/digiconvent/d9t/meta/package_databases"
	"github.com/digiconvent/d9t/meta/services"
	"github.com/digiconvent/embed_env"
	"github.com/digiconvent/install_on_debian"
	"github.com/digiconvent/migrate_packages/.test/log"
)

const name = "digiconvent"

func main() {
	meta.Initialise(os.Args)
	bin := install_on_debian.NewBinary(name)

	switch *flags.Mode {
	case "install":
		preset := *flags.EnvPreset
		var err error
		if bin.IsInstalled() && !*flags.Force {
			panic("cannot install when already installed, use --force to install anyway")
		}
		if bin.IsInstalled() {
			preset, err = embed_env.ReadEmbeddedData("/home/digiconvent/main")
			if err != nil {
				panic("could not read presets from installed binary /home/digiconvent/main: " + err.Error())
			}
			environment.Load(preset)
		} else {
			environment.Env.Prompt(preset)
		}
		err = environment.Save()
		if err != nil {
			panic(err)
		}

		// execute the acme protocol to get https
		err = acme.ExecuteAcmeProtocol("https://acme-v02.api.letsencrypt.org/directory", environment.Env)
		if err != nil {
			panic(err)
		}

		_, err = bin.Install()
		if err != nil {
			panic(err)
		}

		fmt.Println("digiconvent is installed, visit", environment.Env.Domain, "and login with")
		fmt.Println("\n\te-mailaddress:", environment.Env.FirstUser)
		fmt.Println("\tpassword:     ", environment.Env.TelegramBotToken+"\n")
	case "serve":
		log.Info("Serving")
		migrate, err := package_databases.MigrateDatabasesFrom("pkg")
		if err != nil {
			panic(err)
		}
		databases, err := migrate.To("/home/digiconvent/data")
		if err != nil {
			panic(err)
		}

		_, err = services.Initialise(databases)
		if err != nil {
			panic(err)
		}

		server := api_engine.SetupServer()
		server.ListenAndServe()
	case "env":
		installedBinary := "/home/digiconvent/main"
		installedEnv := &environment.EnvVars{}
		if _, err := os.Stat(installedBinary); err == nil {
			fmt.Println("installed version: " + installedEnv.InstalledVersion)
		} else {
			fmt.Println("There is no installed binary")
		}
	}
}
