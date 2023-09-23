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

	var httpClient http.Client
	if cirrusToken == "" {
		// based on cirrus: https://cirrus-ci.org/api/#authorization
		// if you only need read access to public repositories of your organization you can skip this step and don't provide Authorization header.
		httpClient = http.Client{
			Transport: &anonymousTransport{
				wrapped: http.DefaultTransport,
			},
		}

	} else {
		httpClient = http.Client{
			Transport: &authedTransport{
				key:     cirrusToken,
				wrapped: http.DefaultTransport,
			},
		}
	}

	graphqlClient := graphql.NewClient("https://api.cirrus-ci.com/graphql", &httpClient)

	return graphqlClient
}

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

type anonymousTransport struct {
	wrapped http.RoundTripper
}

func (t *anonymousTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.wrapped.RoundTrip(req)
}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
