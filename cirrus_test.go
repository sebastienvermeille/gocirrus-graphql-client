package cirrus

import (
	"context"
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	emptyToken := ""
	graphqlClient := NewClient(emptyToken)
	var response *getRepositoryBranchStatusResponse
	var err error
	response, err = getRepositoryBranchStatus(context.Background(), graphqlClient, "github", "sebastienvermeille", "gocirrus-graphql-client", "master")
	if err != nil {
		return
	}
	fmt.Println("last build status: ", response.OwnerRepository.Builds.Edges[0].Node.Status)
}
