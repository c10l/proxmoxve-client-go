package api2

import (
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
