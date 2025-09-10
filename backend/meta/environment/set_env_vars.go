package environment

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/go-telegram/bot"
)

// this needs manual testing when something is changed
func (envVars *EnvVars) Prompt() {
	// this challenge token will be served and a request to the mux server will be made to verify the domain
	challengeToken := bot.RandomString(40)

	mux := http.NewServeMux()

	mux.HandleFunc("/challenge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, challengeToken)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
	defer srv.Shutdown(context.Background())

	time.Sleep(3 * time.Second)

	envVars.Domain = readInput("fully qualified domain name", envVars.Domain, func(domain string) bool {
		if strings.HasPrefix(domain, "http://") {
			return false
		}
		if strings.HasPrefix(domain, "https://") {
			return false
		}
		if strings.HasPrefix(domain, "www.") {
			return false
		}

		errMsg := fmt.Sprint("Are you sure that ", domain, " is correct? Did you setup the a-record for ", domain, "?")
		resp, err := http.Get("http://" + domain + ":8080/challenge")
		if err != nil {
			fmt.Println(errMsg)
			return false
		}

		defer func() {
			err := resp.Body.Close()
			if err != nil {
				fmt.Println(err.Error())
			}
		}()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(errMsg)
			return false
		}

		if string(body) == challengeToken {
			return true
		} else {
			fmt.Println(errMsg)
			return false
		}
	})
	envVars.FirstUser = readInput("your e-mailaddress", envVars.FirstUser, func(email string) bool {
		_, err := mail.ParseAddress(email)
		return err == nil
	})
	envVars.TelegramBotToken = readInput("a telegram bot token", envVars.TelegramBotToken, func(token string) bool {
		bot, err := bot.New(token)
		if err != nil {
			fmt.Println("Invalid bot token")
			return false
		}
		user, err := bot.GetMe(context.Background())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(user.FirstName + " connected to d9t")
		}
		return true
	})
}

func readInput(question string, defaultAnswer string, validate func(string) bool) string {
	var d string = ""
	if defaultAnswer != "" {
		d = " (default=" + defaultAnswer + ")"
	}
	fmt.Print(question, d, ":")

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
		input = strings.TrimSpace(input)
	} else {
		fmt.Println("Failed to read input.")
	}

	if input == "" && defaultAnswer != "" {
		input = defaultAnswer
	}

	if validate == nil || validate(input) {
		return input
	}

	return readInput(question, defaultAnswer, validate)
}
