// Code generated by Kitex v0.4.2. DO NOT EDIT.

package reviewsservice

import (
	"context"
	reviews "github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ReviewProduct(ctx context.Context, req *reviews.ReviewReq, callOptions ...callopt.Option) (r *reviews.ReviewResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kReviewsServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kReviewsServiceClient struct {
	*kClient
}

func (p *kReviewsServiceClient) ReviewProduct(ctx context.Context, req *reviews.ReviewReq, callOptions ...callopt.Option) (r *reviews.ReviewResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReviewProduct(ctx, req)
}