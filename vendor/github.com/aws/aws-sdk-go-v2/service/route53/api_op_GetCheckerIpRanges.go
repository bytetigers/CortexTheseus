// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// GetCheckerIpRanges still works, but we recommend that you download
// ip-ranges.json, which includes IP address ranges for all Amazon Web Services
// services. For more information, see [IP Address Ranges of Amazon Route 53 Servers]in the Amazon Route 53 Developer Guide.
//
// [IP Address Ranges of Amazon Route 53 Servers]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/route-53-ip-addresses.html
func (c *Client) GetCheckerIpRanges(ctx context.Context, params *GetCheckerIpRangesInput, optFns ...func(*Options)) (*GetCheckerIpRangesOutput, error) {
	if params == nil {
		params = &GetCheckerIpRangesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCheckerIpRanges", params, optFns, c.addOperationGetCheckerIpRangesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCheckerIpRangesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Empty request.
type GetCheckerIpRangesInput struct {
	noSmithyDocumentSerde
}

// A complex type that contains the CheckerIpRanges element.
type GetCheckerIpRangesOutput struct {

	// A complex type that contains sorted list of IP ranges in CIDR format for Amazon
	// Route 53 health checkers.
	//
	// This member is required.
	CheckerIpRanges []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCheckerIpRangesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCheckerIpRanges{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCheckerIpRanges{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCheckerIpRanges"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCheckerIpRanges(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetCheckerIpRanges(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCheckerIpRanges",
	}
}
