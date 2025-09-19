package environment

import (
	"os"

	"github.com/digiconvent/embed_env"
)

type EnvVars struct {
	Domain           string `name:"domain"`             // for obvious reasons, a domain is required
	FirstUser        string `name:"email"`              // main e-mailaddress, for LE and recovery purposes
	TelegramBotToken string `name:"telegram_bot_token"` // this can be edited during runtime, is required for recovery purposes
	TlsCertificate   string `name:"le_cert"`            // this certificate is for TLS
	TlsPrivateKey    string `name:"le_cert_pk"`         // this is the private key for TLS certificates
	TlsAccountId     string `name:"le_account"`         // this is a reference to the account of letsencrypt
	TlsAccountPk     string `name:"le_account_pk"`      // this is the private key for the account
	InstalledVersion string `name:"version"`            // version that was last migrated, this is updated when new migrations run
	JwtPk            string `name:"jwt_private_key"`    // private key for jwts
}

var BinaryVersion string = "-1.-1.-1" // this means dev

var Env *EnvVars = &EnvVars{}

func Load(preset string) error {
	thisBinary, err := os.Executable()
	if err != nil {
		return err
	}
	err = embed_env.ReadFromBinary(thisBinary, Env, preset)
	if err != nil {
		return err
	}
	if Env.InstalledVersion == "" {
		Env.InstalledVersion = "-1.-1.-1"
	}
	return nil
}

func Save() error {
	return embed_env.WriteToBinary(Env)
}

func FromString(s string) *EnvVars {
	env := &EnvVars{}
	embed_env.ReadFromBinary("", env, s)
	return env
}
