package google

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
)

type serviceAccount struct {
	Type         string `json:"type"`
	ProjectID    string `json:"project_id"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
	AuthURI      string `json:"auth_uri"`
	TokenURI     string `json:"token_uri"`
}

func Client() *http.Client {

	f, err := readServiceAccountJSONFile()
	if err != nil {
		log.Fatalf("Unable to read service account json: %v", err)
	}

	sa := serviceAccount{}
	json.Unmarshal(f, &sa)

	conf := &jwt.Config{
		Email:        sa.ClientEmail,
		PrivateKey:   []byte(sa.PrivateKey),
		PrivateKeyID: sa.PrivateKeyID,
		TokenURL:     sa.TokenURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}
	return conf.Client(oauth2.NoContext)
}

func readServiceAccountJSONFile() ([]byte, error) {
	return ioutil.ReadFile(os.Getenv("SERVICE_ACCOUNT_JSON"))
}
