// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	trafficPolicyInstanceOperationTimeout = 4 * time.Minute
)

func waitTrafficPolicyInstanceStateCreated(ctx context.Context, conn *route53.Route53, id string) (*route53.TrafficPolicyInstance, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{TrafficPolicyInstanceStateCreating},
		Target:  []string{TrafficPolicyInstanceStateApplied},
		Refresh: statusTrafficPolicyInstanceState(ctx, conn, id),
		Timeout: trafficPolicyInstanceOperationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*route53.TrafficPolicyInstance); ok {
		if state := aws.StringValue(output.State); state == TrafficPolicyInstanceStateFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.Message)))
		}

		return output, err
	}

	return nil, err
}

func waitTrafficPolicyInstanceStateDeleted(ctx context.Context, conn *route53.Route53, id string) (*route53.TrafficPolicyInstance, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{TrafficPolicyInstanceStateDeleting},
		Target:  []string{},
		Refresh: statusTrafficPolicyInstanceState(ctx, conn, id),
		Timeout: trafficPolicyInstanceOperationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*route53.TrafficPolicyInstance); ok {
		if state := aws.StringValue(output.State); state == TrafficPolicyInstanceStateFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.Message)))
		}

		return output, err
	}

	return nil, err
}

func waitTrafficPolicyInstanceStateUpdated(ctx context.Context, conn *route53.Route53, id string) (*route53.TrafficPolicyInstance, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{TrafficPolicyInstanceStateUpdating},
		Target:  []string{TrafficPolicyInstanceStateApplied},
		Refresh: statusTrafficPolicyInstanceState(ctx, conn, id),
		Timeout: trafficPolicyInstanceOperationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*route53.TrafficPolicyInstance); ok {
		if state := aws.StringValue(output.State); state == TrafficPolicyInstanceStateFailed {
			tfresource.SetLastError(err, errors.New(aws.StringValue(output.Message)))
		}

		return output, err
	}

	return nil, err
}
