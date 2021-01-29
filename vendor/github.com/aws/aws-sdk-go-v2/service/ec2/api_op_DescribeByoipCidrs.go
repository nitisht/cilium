// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the IP address ranges that were specified in calls to
// ProvisionByoipCidr. To describe the address pools that were created when you
// provisioned the address ranges, use DescribePublicIpv4Pools or
// DescribeIpv6Pools.
func (c *Client) DescribeByoipCidrs(ctx context.Context, params *DescribeByoipCidrsInput, optFns ...func(*Options)) (*DescribeByoipCidrsOutput, error) {
	if params == nil {
		params = &DescribeByoipCidrsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeByoipCidrs", params, optFns, addOperationDescribeByoipCidrsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeByoipCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeByoipCidrsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// This member is required.
	MaxResults int32

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The token for the next page of results.
	NextToken *string
}

type DescribeByoipCidrsOutput struct {

	// Information about your address ranges.
	ByoipCidrs []types.ByoipCidr

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeByoipCidrsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeByoipCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeByoipCidrs{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeByoipCidrsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeByoipCidrs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeByoipCidrsAPIClient is a client that implements the DescribeByoipCidrs
// operation.
type DescribeByoipCidrsAPIClient interface {
	DescribeByoipCidrs(context.Context, *DescribeByoipCidrsInput, ...func(*Options)) (*DescribeByoipCidrsOutput, error)
}

var _ DescribeByoipCidrsAPIClient = (*Client)(nil)

// DescribeByoipCidrsPaginatorOptions is the paginator options for
// DescribeByoipCidrs
type DescribeByoipCidrsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeByoipCidrsPaginator is a paginator for DescribeByoipCidrs
type DescribeByoipCidrsPaginator struct {
	options   DescribeByoipCidrsPaginatorOptions
	client    DescribeByoipCidrsAPIClient
	params    *DescribeByoipCidrsInput
	nextToken *string
	firstPage bool
}

// NewDescribeByoipCidrsPaginator returns a new DescribeByoipCidrsPaginator
func NewDescribeByoipCidrsPaginator(client DescribeByoipCidrsAPIClient, params *DescribeByoipCidrsInput, optFns ...func(*DescribeByoipCidrsPaginatorOptions)) *DescribeByoipCidrsPaginator {
	options := DescribeByoipCidrsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeByoipCidrsInput{}
	}

	return &DescribeByoipCidrsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeByoipCidrsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeByoipCidrs page.
func (p *DescribeByoipCidrsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeByoipCidrsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeByoipCidrs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeByoipCidrs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeByoipCidrs",
	}
}
