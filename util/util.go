package util

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/google/go-github/v45/github"
	"github.com/satori/uuid"
	"golang.org/x/oauth2"
)

type githubClient struct {
	Client *github.Client
	Ctx    context.Context
}

var GitHubClient = &githubClient{}

func init() {
	GitHubClient.Ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("github_token")},
	)
	tc := oauth2.NewClient(GitHubClient.Ctx, ts)
	GitHubClient.Client = github.NewClient(tc)
}

func UploadGithub(message string, filename string, img []byte) (downloadURL string, err error) {
	contentResp, resp, err := GitHubClient.Client.Repositories.CreateFile(GitHubClient.Ctx, "imgbjs", "data1", filename, &github.RepositoryContentFileOptions{
		Message: &message,
		Content: img,
	})
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New(http.StatusText(resp.StatusCode))
	}
	return contentResp.Content.GetDownloadURL(), nil
}

func NewUuid() string {
	return uuid.NewV4().String()
}
