package precheck_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/controllers/cluster-mgr/precheck"
	corev1 "k8s.io/api/core/v1"
)

func TestDatabasePreChecker(t *testing.T) {
	db := &wutongv1alpha1.Database{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "foo",
		Password: "bar",
		Name:     "foobar",
	}

	preChecker := precheck.NewDatabasePrechecker(wutongv1alpha1.WutongClusterConditionTypeDatabaseRegion, db)

	condition := preChecker.Check()

	assert.Equal(t, wutongv1alpha1.WutongClusterConditionType(wutongv1alpha1.WutongClusterConditionTypeDatabaseRegion), condition.Type)
	assert.Equal(t, corev1.ConditionFalse, condition.Status)
	assert.Equal(t, "DatabaseFailed", condition.Reason)
}
