package acme

import (
	"github.com/digiconvent/d9t/meta/environment"
	"github.com/digiconvent/migrate_packages/.test/log"
	"github.com/digiconvent/setup_acme"
)

func ExecuteAcmeProtocol(directoryUrl string, env *environment.EnvVars) error {
	acmeClient := setup_acme.AcmeClient{
		DirectoryUrl: directoryUrl,
		InitData: &setup_acme.InitData{
			Domain:       env.Domain,
			Emailaddress: env.FirstUser,
			Organisation: "DigiConvent",
		},
		RefreshData: &setup_acme.RefreshData{},
	}
	err := acmeClient.Do()
	if err != nil {
		log.Error(err)
		return err
	}

	env.TlsAccountId = acmeClient.RefreshData.Kid
	env.TlsAccountPk = acmeClient.InitData.AccountPrivateKey
	env.TlsPrivateKey = acmeClient.InitData.DomainPrivateKey
	env.TlsCertificate = acmeClient.RefreshData.Certificate

	return environment.Save()
}
