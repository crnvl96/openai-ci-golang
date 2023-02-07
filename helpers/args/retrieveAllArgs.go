package args

import (
	"fmt"
	"os"

	i "github.com/crnvl96/openai-ci-golang/helpers/args/individualArgs"
)

type Args struct {
	AIAPIKey string
	RepositoryOwner string
	RepositoryName string
	PullRequestNumber int
	GithubToken string
}

func GetAllArgs() Args {
	args := Args{}

	APIKey, err := i.RetrieveAIAPIKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repositoryOwner, err := i.RetrieveRepositoryOwner()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repositoryName, err := i.RetrieveRepositoryName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pullRequestNumber, err := i.RetrievePullRequestNumber()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	githubToken, err := i.RetrieveGithubToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args.AIAPIKey = APIKey
	args.RepositoryOwner = repositoryOwner
	args.RepositoryName = repositoryName
	args.PullRequestNumber = pullRequestNumber
	args.GithubToken = githubToken

	return args
}