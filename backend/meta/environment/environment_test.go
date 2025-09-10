package environment_test

import (
	"os"
	"testing"

	"github.com/digiconvent/d9t/meta"
	"github.com/digiconvent/d9t/meta/environment"
)

func TestEnvironmentLoadingAndSaving(t *testing.T) {
	// this should be loaded later
	env := environment.Env
	env.Domain = "somedomain.de"
	env.TelegramBotToken = "awdawd"
	environment.Save()

	env.Domain = "wrongdomain.de"
	env.TelegramBotToken = "wrongtoken"

	err := meta.Initialise(os.Args)
	if err == nil {
		t.Fatal("expected Initialise to throw an error")
	}

	err = meta.Initialise([]string{"this is the call for the binary, the following strings are the args", "-mode=install"})
	if err != nil {
		t.Fatal("expected Initialise not to throw an error, instead got", err)
	}

	if err != nil {
		t.Fatal("did not expect environment.Load()")
	}

	// env should now contain the values that were saved
	if env.Domain != "somedomain.de" || env.TelegramBotToken != "awdawd" {
		t.Fatal("expected env to be loaded with previous data")
	}
}
