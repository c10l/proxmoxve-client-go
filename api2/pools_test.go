package api2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPostPoolAndGetPools(t *testing.T) {
	poolID := rand.String(10)
	comment := rand.String(20)
	assert.NoError(t, testClient.PostPool(poolID, comment))

	pools, getPoolsError := testClient.GetPools()
	assert.NoError(t, getPoolsError)
	assert.Contains(t, *pools, Pool{PoolID: poolID, Comment: comment})
}

func TestGetPool(t *testing.T) {
	poolID := rand.String(10)
	comment := rand.String(20)
	assert.NoError(t, testClient.PostPool(poolID, comment))

	expected := Pool{
		PoolID:  poolID,
		Comment: comment,
		Members: []any{},
	}
	actual, getPoolError := testClient.GetPool(poolID)
	assert.NoError(t, getPoolError)
	assert.Equal(t, expected, *actual)
}

func TestDeletePool(t *testing.T) {
	poolID := rand.String(10)
	assert.NoError(t, testClient.PostPool(poolID, ""))
	assert.NoError(t, testClient.DeletePool(poolID))
	_, getPoolError := testClient.GetPool(poolID)
	assert.ErrorContains(t, getPoolError, fmt.Sprintf("500 pool '%s' does not exist", poolID))
}
