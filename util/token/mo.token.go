package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func GenerateToken(r *http.Request) (string, error) {
	// convert request to json
	fmt.Println("request = ", r.URL.Query())
	requestQuery, err := json.Marshal(r.URL.Query())
	if err != nil {
		return "", err
	}

	register := exec.Command("./registermo")
	register.Stdin = strings.NewReader(string(requestQuery))

	var out bytes.Buffer
	register.Stdout = &out
	err := register.Run()
	if err != nil {
		fmt.Println(err)
		return "empty"
	}

	return out.String()
}
