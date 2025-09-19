package acme_test

import (
	"fmt"
	"testing"

	"github.com/digiconvent/d9t/meta/acme"
	"github.com/digiconvent/d9t/meta/environment"
)

func TestAcmeProtocol(t *testing.T) {
	domain := "digiconvent.de"
	// this can only be tested on a server with a records for domain
	env := environment.FromString("domain=" + domain + "&email=info@" + domain)

	fmt.Println(env.Domain)
	fmt.Println(env.FirstUser)
	err := acme.ExecuteAcmeProtocol("https://acme-staging-v02.api.letsencrypt.org/directory", env)

	if err != nil {
		t.Fatal(err)
	}
}
