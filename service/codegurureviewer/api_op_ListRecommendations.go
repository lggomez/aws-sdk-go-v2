// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListRecommendationsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the code review to describe.
	//
	// CodeReviewArn is a required field
	CodeReviewArn *string `location:"uri" locationName:"CodeReviewArn" min:"1" type:"string" required:"true"`

	// The maximum number of results that are returned per call. The default is
	// 100.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// Pagination token.
	NextToken *string `location:"querystring" locationName:"NextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListRecommendationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRecommendationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRecommendationsInput"}

	if s.CodeReviewArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CodeReviewArn"))
	}
	if s.CodeReviewArn != nil && len(*s.CodeReviewArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CodeReviewArn", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRecommendationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CodeReviewArn != nil {
		v := *s.CodeReviewArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CodeReviewArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListRecommendationsOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token.
	NextToken *string `min:"1" type:"string"`

	// List of recommendations for the requested code review.
	RecommendationSummaries []RecommendationSummary `type:"list"`
}

// String returns the string representation
func (s ListRecommendationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRecommendationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RecommendationSummaries != nil {
		v := s.RecommendationSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RecommendationSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListRecommendations = "ListRecommendations"

// ListRecommendationsRequest returns a request value for making API operation for
// Amazon CodeGuru Reviewer.
//
// Returns the list of all recommendations for a completed code review.
//
//    // Example sending a request using ListRecommendationsRequest.
//    req := client.ListRecommendationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguru-reviewer-2019-09-19/ListRecommendations
func (c *Client) ListRecommendationsRequest(input *ListRecommendationsInput) ListRecommendationsRequest {
	op := &aws.Operation{
		Name:       opListRecommendations,
		HTTPMethod: "GET",
		HTTPPath:   "/codereviews/{CodeReviewArn}/Recommendations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRecommendationsInput{}
	}

	req := c.newRequest(op, input, &ListRecommendationsOutput{})
	return ListRecommendationsRequest{Request: req, Input: input, Copy: c.ListRecommendationsRequest}
}

// ListRecommendationsRequest is the request type for the
// ListRecommendations API operation.
type ListRecommendationsRequest struct {
	*aws.Request
	Input *ListRecommendationsInput
	Copy  func(*ListRecommendationsInput) ListRecommendationsRequest
}

// Send marshals and sends the ListRecommendations API request.
func (r ListRecommendationsRequest) Send(ctx context.Context) (*ListRecommendationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRecommendationsResponse{
		ListRecommendationsOutput: r.Request.Data.(*ListRecommendationsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRecommendationsRequestPaginator returns a paginator for ListRecommendations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRecommendationsRequest(input)
//   p := codegurureviewer.NewListRecommendationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRecommendationsPaginator(req ListRecommendationsRequest) ListRecommendationsPaginator {
	return ListRecommendationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRecommendationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListRecommendationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRecommendationsPaginator struct {
	aws.Pager
}

func (p *ListRecommendationsPaginator) CurrentPage() *ListRecommendationsOutput {
	return p.Pager.CurrentPage().(*ListRecommendationsOutput)
}

// ListRecommendationsResponse is the response type for the
// ListRecommendations API operation.
type ListRecommendationsResponse struct {
	*ListRecommendationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRecommendations request.
func (r *ListRecommendationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
