// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For the specified group in the specified identity store, returns the list of all
// GroupMembership objects and returns results in paginated form.
func (c *Client) ListGroupMemberships(ctx context.Context, params *ListGroupMembershipsInput, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) {
	if params == nil {
		params = &ListGroupMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMemberships", params, optFns, c.addOperationListGroupMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembershipsInput struct {

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The maximum number of results to be returned per request. This parameter is used
	// in the ListUsers and ListGroups requests to specify how many results to return
	// in one page. The length limit is 50 characters.
	MaxResults *int32

	// The pagination token used for the ListUsers, ListGroups and ListGroupMemberships
	// API operations. This value is generated by the identity store service. It is
	// returned in the API response if the total results are more than the size of one
	// page. This token is also returned when it is used in the API request to search
	// for the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupMembershipsOutput struct {

	// A list of GroupMembership objects in the group.
	//
	// This member is required.
	GroupMemberships []types.GroupMembership

	// The pagination token used for the ListUsers, ListGroups and ListGroupMemberships
	// API operations. This value is generated by the identity store service. It is
	// returned in the API response if the total results are more than the size of one
	// page. This token is also returned when it is used in the API request to search
	// for the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroupMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroupMemberships{}, middleware.After)
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
	if err = addOpListGroupMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupMemberships(options.Region), middleware.Before); err != nil {
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

// ListGroupMembershipsAPIClient is a client that implements the
// ListGroupMemberships operation.
type ListGroupMembershipsAPIClient interface {
	ListGroupMemberships(context.Context, *ListGroupMembershipsInput, ...func(*Options)) (*ListGroupMembershipsOutput, error)
}

var _ ListGroupMembershipsAPIClient = (*Client)(nil)

// ListGroupMembershipsPaginatorOptions is the paginator options for
// ListGroupMemberships
type ListGroupMembershipsPaginatorOptions struct {
	// The maximum number of results to be returned per request. This parameter is used
	// in the ListUsers and ListGroups requests to specify how many results to return
	// in one page. The length limit is 50 characters.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupMembershipsPaginator is a paginator for ListGroupMemberships
type ListGroupMembershipsPaginator struct {
	options   ListGroupMembershipsPaginatorOptions
	client    ListGroupMembershipsAPIClient
	params    *ListGroupMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListGroupMembershipsPaginator returns a new ListGroupMembershipsPaginator
func NewListGroupMembershipsPaginator(client ListGroupMembershipsAPIClient, params *ListGroupMembershipsInput, optFns ...func(*ListGroupMembershipsPaginatorOptions)) *ListGroupMembershipsPaginator {
	if params == nil {
		params = &ListGroupMembershipsInput{}
	}

	options := ListGroupMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupMemberships page.
func (p *ListGroupMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListGroupMemberships(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListGroupMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "ListGroupMemberships",
	}
}
