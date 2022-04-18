package api2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestPostPoolAndGetPoolList(t *testing.T) {
	poolID := rand.String(10)
	comment := rand.String(20)
	assert.NoError(t, testClient.PostPool(poolID, comment))

	poolList, getPoolListError := testClient.GetPoolList()
	assert.NoError(t, getPoolListError)
	assert.Contains(t, *poolList, Pool{PoolID: poolID, Comment: comment})
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

func TestPutPoolModifyComment(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	assert.NoError(t, testClient.PostPool(poolID, ""))
	assert.NoError(t, testClient.PutPool(poolID, &expectedComment, nil, nil, false))
	pool, _ := testClient.GetPool(poolID)
	assert.Equal(t, expectedComment, pool.Comment)
}

func TestPutPoolAddAndDeleteStorage(t *testing.T) {
	poolID := rand.String(10)
	expectedStorage := "local-lvm"
	assert.NoError(t, testClient.PostPool(poolID, ""))
	assert.NoError(t, testClient.PutPool(poolID, nil, &[]string{expectedStorage}, nil, false))
	pool, _ := testClient.GetPool(poolID)
	members := pool.Members[0].(map[string]any)
	assert.Equal(t, members["storage"], expectedStorage)

	assert.NoError(t, testClient.PutPool(poolID, nil, &[]string{expectedStorage}, nil, true))
	pool, _ = testClient.GetPool(poolID)
	assert.Len(t, pool.Members, 0)
}

func TestPutPoolAddAndDeleteVMs(t *testing.T) {
	// TODO: Needs implementing adding VMs/CTs
}
