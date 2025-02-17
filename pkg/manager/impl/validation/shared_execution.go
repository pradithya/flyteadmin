package validation

import (
	"context"
	"fmt"

	"github.com/flyteorg/flyteadmin/pkg/common"
	"github.com/flyteorg/flyteadmin/pkg/errors"
	"github.com/flyteorg/flyteadmin/pkg/repositories"
	repoInterfaces "github.com/flyteorg/flyteadmin/pkg/repositories/interfaces"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flytestdlib/logger"
)

// ValidateClusterForExecutionID validates that the execution denoted by executionId is recorded as executing on `cluster`.
func ValidateClusterForExecutionID(ctx context.Context, db repositories.RepositoryInterface, executionID *core.WorkflowExecutionIdentifier, cluster string) error {
	workflowExecution, err := db.ExecutionRepo().Get(ctx, repoInterfaces.Identifier{
		Project: executionID.Project,
		Domain:  executionID.Domain,
		Name:    executionID.Name,
	})
	if err != nil {
		logger.Debugf(ctx, "Failed to find existing execution with id [%+v] with err: %v", executionID, err)
		return err
	}
	return ValidateCluster(ctx, workflowExecution.Cluster, cluster)
}

// ValidateClusterForExecution validates that the execution is recorded as executing on `cluster`.
func ValidateCluster(ctx context.Context, recordedCluster, cluster string) error {
	// DefaultProducerID is used in older versions of propeller which hard code this producer id.
	// See https://github.com/flyteorg/flytepropeller/blob/eaf084934de5d630cd4c11aae15ecae780cc787e/pkg/controller/nodes/task/transformer.go#L114
	if len(cluster) == 0 || cluster == common.DefaultProducerID {
		return nil
	}
	if recordedCluster != cluster {
		errorMsg := fmt.Sprintf("Cluster/producer from event [%s] does not match existing workflow execution cluster: [%s]",
			recordedCluster, cluster)
		return errors.NewIncompatibleClusterError(ctx, errorMsg, recordedCluster)
	}
	return nil
}
