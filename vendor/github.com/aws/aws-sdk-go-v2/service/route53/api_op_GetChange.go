// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the current status of a change batch request. The status is one of the
// following values:
//
// * PENDING indicates that the changes in this request have not
// propagated to all Amazon Route 53 DNS servers. This is the initial status of all
// change batch requests.
//
// * INSYNC indicates that the changes have propagated to
// all Route 53 DNS servers.
func (c *Client) GetChange(ctx context.Context, params *GetChangeInput, optFns ...func(*Options)) (*GetChangeOutput, error) {
	if params == nil {
		params = &GetChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChange", params, optFns, c.addOperationGetChangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for a GetChange request.
type GetChangeInput struct {

	// The ID of the change batch request. The value that you specify here is the value
	// that ChangeResourceRecordSets returned in the Id element when you submitted the
	// request.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// A complex type that contains the ChangeInfo element.
type GetChangeOutput struct {

	// A complex type that contains information about the specified change batch.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetChange{}, middleware.After)
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
	if err = addOpGetChangeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChange(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetChangeAPIClient is a client that implements the GetChange operation.
type GetChangeAPIClient interface {
	GetChange(context.Context, *GetChangeInput, ...func(*Options)) (*GetChangeOutput, error)
}

var _ GetChangeAPIClient = (*Client)(nil)

// ResourceRecordSetsChangedWaiterOptions are waiter options for
// ResourceRecordSetsChangedWaiter
type ResourceRecordSetsChangedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ResourceRecordSetsChangedWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ResourceRecordSetsChangedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetChangeInput, *GetChangeOutput, error) (bool, error)
}

// ResourceRecordSetsChangedWaiter defines the waiters for
// ResourceRecordSetsChanged
type ResourceRecordSetsChangedWaiter struct {
	client GetChangeAPIClient

	options ResourceRecordSetsChangedWaiterOptions
}

// NewResourceRecordSetsChangedWaiter constructs a ResourceRecordSetsChangedWaiter.
func NewResourceRecordSetsChangedWaiter(client GetChangeAPIClient, optFns ...func(*ResourceRecordSetsChangedWaiterOptions)) *ResourceRecordSetsChangedWaiter {
	options := ResourceRecordSetsChangedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = resourceRecordSetsChangedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ResourceRecordSetsChangedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ResourceRecordSetsChanged waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ResourceRecordSetsChangedWaiter) Wait(ctx context.Context, params *GetChangeInput, maxWaitDur time.Duration, optFns ...func(*ResourceRecordSetsChangedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ResourceRecordSetsChanged waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ResourceRecordSetsChangedWaiter) WaitForOutput(ctx context.Context, params *GetChangeInput, maxWaitDur time.Duration, optFns ...func(*ResourceRecordSetsChangedWaiterOptions)) (*GetChangeOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetChange(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ResourceRecordSetsChanged waiter")
}

func resourceRecordSetsChangedStateRetryable(ctx context.Context, input *GetChangeInput, output *GetChangeOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ChangeInfo.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "INSYNC"
		value, ok := pathValue.(types.ChangeStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ChangeStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetChange(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetChange",
	}
}
