package elsys

import (
	"os"
	"os/user"
)

// EnvUsername returns the username from the OS environment.
func EnvUsername() (string, error) {
	return os.Getenv("USER"), nil
}

// OSUsername returns the username of the current OS user (based on UID).
func OSUsername() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.Username, nil
}
