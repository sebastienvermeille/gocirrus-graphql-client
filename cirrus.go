package cirrus

import (
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"net/http"
	"os"
)

// NewClient returns a new Cirrus API client.
func NewClient(cirrusToken string) graphql.Client {

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if cirrusToken == "" {
		err = fmt.Errorf("must set GITHUB_TOKEN=<github token>")
		return nil
	}

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     cirrusToken,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient("https://api.cirrus-ci.com/graphql", &httpClient)

	return graphqlClient
}
