package environment

import "github.com/digiconvent/embed_env"

type EnvVars struct {
	Domain           string `name:"domain"`             // for obvious reasons, a domain is required
	FirstUser        string `name:"email"`              // main e-mailaddress, for LE and recovery purposes
	TelegramBotToken string `name:"telegram_bot_token"` // this can be edited during runtime, is required for recovery purposes
	TlsCertificate   string `name:"le_cert"`            // this certificate is for TLS
	TlsPrivateKey    string `name:"le_cert_pk"`         // this is the private key for TLS certificates
	TlsAccountId     string `name:"le_account"`         // this is a reference to the account of letsencrypt
	TlsAccountPk     string `name:"le_account_pk"`      // this is the private key for the account
}

var Env *EnvVars = &EnvVars{}

func Load() error {
	return embed_env.ReadFromBinary(Env)
}

func Save() error {
	return embed_env.WriteToBinary(Env)
}
