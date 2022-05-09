package api2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPoolGet(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	err := testClient.PostPool(poolID, expectedComment)
	assert.NoError(t, err)

	pool, err := testClient.GetPool(poolID)
	assert.NoError(t, err)
	assert.Equal(t, poolID, pool.PoolID)
	assert.Equal(t, expectedComment, pool.Comment)
}

func TestPoolPut(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)

	err := testClient.PostPool(poolID, expectedComment)
	assert.NoError(t, err)

	err = testClient.PutPool(poolID, &expectedComment, nil, nil, false)
	assert.NoError(t, err)

	pool, err := testClient.GetPool(poolID)
	assert.NoError(t, err)
	assert.Equal(t, expectedComment, pool.Comment)
}

func TestPoolDelete(t *testing.T) {
	poolID := rand.String(10)

	err := testClient.PostPool(poolID, "")
	assert.NoError(t, err)

	_, err = testClient.GetPool(poolID)
	assert.NoError(t, err)

	err = testClient.DeletePool(poolID)
	assert.NoError(t, err)

	_, err = testClient.GetPool(poolID)
	assert.ErrorContains(t, err, fmt.Sprintf("500 pool '%s' does not exist", poolID))
}
