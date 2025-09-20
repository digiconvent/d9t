package acme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/digiconvent/d9t/meta/acme"
	"github.com/digiconvent/d9t/meta/environment"
)

func TestAcmeProtocol(t *testing.T) {
	if os.Getenv("ACME_TEST") == "" {
		t.Skip("Skipping ACME test - set ACME_TEST=1 to run on server with proper DNS")
	}

	domain := "digiconvent.de"
	env := environment.FromString("domain=" + domain + "&email=info@" + domain)

	fmt.Println(env.Domain)
	fmt.Println(env.FirstUser)
	err := acme.ExecuteAcmeProtocol("https://acme-staging-v02.api.letsencrypt.org/directory", env)

	if err != nil {
		t.Fatal(err)
	}
}
