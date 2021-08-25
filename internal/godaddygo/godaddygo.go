package godaddygo
import (
    "os/user"
    "encoding/json"
    "io/ioutil"
    "github.com/oze4/godaddygo"
)

type Credentials struct {
    Key string `json:"key"`
    Secret string `json:"secret"`
}

// getCredentials reads the credentials files located in $HOME/.config/godaddy/credentials.json and returns a Credentials
func getCredentials() (Credentials, error) {
    var creds Credentials
    // Get current user in order to get credentials
    user, err := user.Current()

    if err != nil {
        return creds, err
    }

    data, err := ioutil.ReadFile(user.HomeDir + "/.config/godaddy/credentials.json")

    if err != nil {
        return creds, err
    }

    err = json.Unmarshal(data, &creds)

    if err != nil {
        return creds, err
    }

    return creds, nil
}

func GetGoDaddy() (godaddygo.V1, error) {
    creds, err := getCredentials()
    if err != nil {
        return nil, err
    }

    api, err := godaddygo.NewProduction(creds.Key, creds.Secret)
    if err != nil {
        return nil, err
    }
    godaddy := api.V1()

    return godaddy, nil
}
