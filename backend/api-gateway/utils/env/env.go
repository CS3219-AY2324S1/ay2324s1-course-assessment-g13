package env

import "os"

const (
	GITHUB_CLIENT_ID_ENV_KEY     = "GITHUB_CLIENT_ID"
	GITHUB_CLIENT_SECRET_ENV_KEY = "GITHUB_CLIENT_SECRET"
)

func GetGitHubClientID() string {
	return os.Getenv(GITHUB_CLIENT_ID_ENV_KEY)
}

func GetGitHubClientSecret() string {
	return os.Getenv(GITHUB_CLIENT_SECRET_ENV_KEY)
}
