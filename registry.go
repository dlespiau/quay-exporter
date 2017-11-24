package main

import (
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/dlespiau/quay-exporter/pkg/models"
	"github.com/dlespiau/quay-exporter/pkg/quay"
	"github.com/dlespiau/quay-exporter/pkg/quay/repository"
)

const (
	quayHost   = "quay.io"
	quayPath   = "/"
	quayScheme = "https"
)

// Registry is an Docker image registry.
type Registry struct {
	token  string
	client *quay.Client
}

// NewRegistry creates a new registry.
func NewRegistry() *Registry {
	transport := httptransport.New(quayHost, quayPath, []string{quayScheme})
	client := quay.New(transport, strfmt.Default)

	return &Registry{
		client: client,
	}
}

func (r *Registry) SetBearerToken(token string) {
	r.token = token
}

func (r *Registry) getAuth() runtime.ClientAuthInfoWriter {
	if r.token != "" {
		return httptransport.BearerToken(r.token)
	}
	return nil
}

// ListRepositories returns the list of available Docker image repositories in a
// given namespace.
func (r *Registry) ListRepositories(namespace string) (models.ListReposRepositories, error) {

	params := repository.NewListReposParamsWithTimeout(10 * time.Second)
	params.Namespace = &namespace
	kind := "image"
	params.RepoKind = &kind

	// If we don't have a token, just ask for public images. If we try to ask for
	// public & private images without a token, the API returns an empty list of
	// repositories.
	if r.token == "" {
		public := true
		params.Public = &public
	}

	response, err := r.client.Repository.ListRepos(params, r.getAuth())
	if err != nil {
		return nil, err
	}
	return response.Payload.Repositories, nil
}
