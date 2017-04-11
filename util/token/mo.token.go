package token

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Mo Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func GenerateToken(r *http.Request) (string, error) {
	// convert request to json
	requestQuery, err := json.Marshal(r.URL.Query())
	if err != nil {
		return "", err
	}

	//register := exec.Command(".registermo")
	register := exec.Command("util/token/registermo")
	register.Stdin = strings.NewReader(string(requestQuery))

	var out bytes.Buffer
	register.Stdout = &out
	err = register.Run()
	if err != nil {
		logger.Println(err)
		return "", err
	}

	return out.String(), nil
}

func GenerateTokenString(input string) (string, error) {
	//register := exec.Command("register")
	register := exec.Command("/util/token/registermo")
	register.Stdin = strings.NewReader(input)

	var out bytes.Buffer
	register.Stdout = &out
	err := register.Run()
	if err != nil {
		logger.Println(err)
		return "", err
	}

	return out.String(), nil
}
