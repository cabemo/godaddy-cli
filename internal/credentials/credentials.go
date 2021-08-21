package credentials
import (
    "os/user"
    "encoding/json"
    "io/ioutil"
)

type Credentials struct {
    Key string `json:"key"`
    Secret string `json:"secret"`
}

func GetCredentials() (Credentials, error) {
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
