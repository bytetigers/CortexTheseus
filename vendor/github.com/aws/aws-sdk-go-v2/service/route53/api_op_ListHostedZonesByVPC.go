// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which Amazon Web Services account or Amazon Web Services service
// owns the hosted zones. The HostedZoneOwner structure in the response contains
// one of the following values:
//   - An OwningAccount element, which contains the account number of either the
//     current Amazon Web Services account or another Amazon Web Services account. Some
//     services, such as Cloud Map, create hosted zones using the current account.
//   - An OwningService element, which identifies the Amazon Web Services service
//     that created and owns the hosted zone. For example, if a hosted zone was created
//     by Amazon Elastic File System (Amazon EFS), the value of Owner is
//     efs.amazonaws.com .
//
// When listing private hosted zones, the hosted zone and the Amazon VPC must
// belong to the same partition where the hosted zones were created. A partition is
// a group of Amazon Web Services Regions. Each Amazon Web Services account is
// scoped to one partition. The following are the supported partitions:
//   - aws - Amazon Web Services Regions
//   - aws-cn - China Regions
//   - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see Access Management (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// in the Amazon Web Services General Reference.
func (c *Client) ListHostedZonesByVPC(ctx context.Context, params *ListHostedZonesByVPCInput, optFns ...func(*Options)) (*ListHostedZonesByVPCOutput, error) {
	if params == nil {
		params = &ListHostedZonesByVPCInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostedZonesByVPC", params, optFns, c.addOperationListHostedZonesByVPCMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostedZonesByVPCOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which Amazon Web Services account created the hosted zones.
type ListHostedZonesByVPCInput struct {

	// The ID of the Amazon VPC that you want to list hosted zones for.
	//
	// This member is required.
	VPCId *string

	// For the Amazon VPC that you specified for VPCId , the Amazon Web Services Region
	// that you created the VPC in.
	//
	// This member is required.
	VPCRegion types.VPCRegion

	// (Optional) The maximum number of hosted zones that you want Amazon Route 53 to
	// return. If the specified VPC is associated with more than MaxItems hosted
	// zones, the response includes a NextToken element. NextToken contains an
	// encrypted token that identifies the first hosted zone that Route 53 will return
	// if you submit another request.
	MaxItems *int32

	// If the previous response included a NextToken element, the specified VPC is
	// associated with more hosted zones. To get more hosted zones, submit another
	// ListHostedZonesByVPC request. For the value of NextToken , specify the value of
	// NextToken from the previous response. If the previous response didn't include a
	// NextToken element, there are no more hosted zones to get.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHostedZonesByVPCOutput struct {

	// A list that contains one HostedZoneSummary element for each hosted zone that
	// the specified Amazon VPC is associated with. Each HostedZoneSummary element
	// contains the hosted zone name and ID, and information about who owns the hosted
	// zone.
	//
	// This member is required.
	HostedZoneSummaries []types.HostedZoneSummary

	// The value that you specified for MaxItems in the most recent
	// ListHostedZonesByVPC request.
	//
	// This member is required.
	MaxItems *int32

	// The value that you will use for NextToken in the next ListHostedZonesByVPC
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHostedZonesByVPCMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListHostedZonesByVPC{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListHostedZonesByVPC{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addListHostedZonesByVPCResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListHostedZonesByVPCValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedZonesByVPC(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListHostedZonesByVPC(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListHostedZonesByVPC",
	}
}

type opListHostedZonesByVPCResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListHostedZonesByVPCResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListHostedZonesByVPCResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "route53"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "route53"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("route53")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListHostedZonesByVPCResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListHostedZonesByVPCResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
