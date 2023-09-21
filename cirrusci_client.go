// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package client

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// getViewerResponse is returned by getViewer on success.
type getViewerResponse struct {
	Viewer getViewerViewerUser `json:"viewer"`
}

// GetViewer returns getViewerResponse.Viewer, and is useful for accessing the field via an interface.
func (v *getViewerResponse) GetViewer() getViewerViewerUser { return v.Viewer }

// getViewerViewerUser includes the requested fields of the GraphQL type User.
type getViewerViewerUser struct {
	Id string `json:"id"`
}

// GetId returns getViewerViewerUser.Id, and is useful for accessing the field via an interface.
func (v *getViewerViewerUser) GetId() string { return v.Id }

// The query or mutation executed by getViewer.
const getViewer_Operation = `
query getViewer {
	viewer {
		id
	}
}
`

func getViewer(
	ctx context.Context,
	client graphql.Client,
) (*getViewerResponse, error) {
	req := &graphql.Request{
		OpName: "getViewer",
		Query:  getViewer_Operation,
	}
	var err error

	var data getViewerResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}