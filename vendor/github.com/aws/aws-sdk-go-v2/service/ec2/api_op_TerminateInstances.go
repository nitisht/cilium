// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shuts down the specified instances. This operation is idempotent; if you
// terminate an instance more than once, each call succeeds. If you specify
// multiple instances and the request fails (for example, because of a single
// incorrect instance ID), none of the instances are terminated. Terminated
// instances remain visible after termination (for approximately one hour). By
// default, Amazon EC2 deletes all EBS volumes that were attached when the instance
// launched. Volumes attached after instance launch continue running. You can stop,
// start, and terminate EBS-backed instances. You can only terminate instance
// store-backed instances. What happens to an instance differs if you stop it or
// terminate it. For example, when you stop an instance, the root device and any
// other devices attached to the instance persist. When you terminate an instance,
// any attached EBS volumes with the DeleteOnTermination block device mapping
// parameter set to true are automatically deleted. For more information about the
// differences between stopping and terminating instances, see Instance lifecycle
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information about
// troubleshooting, see Troubleshooting terminating your instance
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesShuttingDown.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) TerminateInstances(ctx context.Context, params *TerminateInstancesInput, optFns ...func(*Options)) (*TerminateInstancesOutput, error) {
	if params == nil {
		params = &TerminateInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateInstances", params, optFns, addOperationTerminateInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateInstancesInput struct {

	// The IDs of the instances. Constraints: Up to 1000 instance IDs. We recommend
	// breaking up this request into smaller batches.
	//
	// This member is required.
	InstanceIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

type TerminateInstancesOutput struct {

	// Information about the terminated instances.
	TerminatingInstances []types.InstanceStateChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTerminateInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpTerminateInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpTerminateInstances{}, middleware.After)
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
	if err = addOpTerminateInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTerminateInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "TerminateInstances",
	}
}
