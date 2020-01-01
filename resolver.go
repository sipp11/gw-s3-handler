package gw_s3_handler

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Upload(ctx context.Context, file graphql.Upload) (*S3File, error) {
	result, err := S3Upload("satang-expensy", file)
	if err != nil {
		return nil, err
	}
	s3File := S3File{Location: result.Location}
	// TODO: update DB as well
	return &s3File, nil
}

func (r *mutationResolver) MultipleUpload(ctx context.Context, appName string, files []*UploadFile, table *string, column *string, requireAbsPath *bool) ([]*S3File, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Files(ctx context.Context) ([]*S3File, error) {
	// TODO: probably fetching from DB to answer
	TestDB()
	return nil, nil
	// panic("not implemented")

}
