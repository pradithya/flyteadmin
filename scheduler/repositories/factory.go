package repositories

import (
	"fmt"

	"github.com/flyteorg/flyteadmin/pkg/repositories/config"
	"github.com/flyteorg/flyteadmin/pkg/repositories/errors"
	"github.com/flyteorg/flyteadmin/scheduler/repositories/interfaces"
	"github.com/flyteorg/flytestdlib/promutils"
)

type RepoConfig int32

const (
	POSTGRES RepoConfig = 0
)

var RepositoryConfigurationName = map[int32]string{
	0: "POSTGRES",
}

// The SchedulerRepoInterface indicates the methods that each Repository must support.
// A Repository indicates a Database which is collection of Tables/models.
// The goal is allow databases to be Plugged in easily.
// This interface contains only scheduler specific models and tables.

type SchedulerRepoInterface interface {
	SchedulableEntityRepo() interfaces.SchedulableEntityRepoInterface
	ScheduleEntitiesSnapshotRepo() interfaces.ScheduleEntitiesSnapShotRepoInterface
}

func GetRepository(repoType RepoConfig, dbConfig config.DbConfig, scope promutils.Scope) SchedulerRepoInterface {
	switch repoType {
	case POSTGRES:
		postgresScope := scope.NewSubScope("postgres")
		db, err := config.OpenDbConnection(config.NewPostgresConfigProvider(dbConfig, postgresScope))
		if err != nil {
			panic(err)
		}
		return NewPostgresRepo(
			db,
			errors.NewPostgresErrorTransformer(postgresScope.NewSubScope("errors")),
			postgresScope.NewSubScope("repositories"))
	default:
		panic(fmt.Sprintf("Invalid repoType %v", repoType))
	}
}
