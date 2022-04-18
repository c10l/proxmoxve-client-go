package api2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestCreatePoolAndRetrievePoolList(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	assert.NoError(t, testClient.CreatePool(poolID, expectedComment))

	actualPoolList, retrievePoolListError := testClient.RetrievePoolList()
	assert.NoError(t, retrievePoolListError)
	assert.Contains(t, *actualPoolList, Pool{PoolID: poolID, Comment: expectedComment})
}

func TestRetrievePool(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	assert.NoError(t, testClient.CreatePool(poolID, expectedComment))

	expectedPool := Pool{
		PoolID:  poolID,
		Comment: expectedComment,
		Members: []any{},
	}
	actualPool, retrievePoolError := testClient.RetrievePool(poolID)
	assert.NoError(t, retrievePoolError)
	assert.Equal(t, expectedPool, *actualPool)
}

func TestDeletePool(t *testing.T) {
	poolID := rand.String(10)
	assert.NoError(t, testClient.CreatePool(poolID, ""))
	assert.NoError(t, testClient.DeletePool(poolID))
	_, retrievePoolError := testClient.RetrievePool(poolID)
	assert.ErrorContains(t, retrievePoolError, fmt.Sprintf("500 pool '%s' does not exist", poolID))
}

func TestUpdatePoolComment(t *testing.T) {
	poolID := rand.String(10)
	expectedComment := rand.String(20)
	assert.NoError(t, testClient.CreatePool(poolID, ""))
	assert.NoError(t, testClient.UpdatePool(poolID, &expectedComment, nil, nil, false))
	actualPool, _ := testClient.RetrievePool(poolID)
	assert.Equal(t, expectedComment, actualPool.Comment)
}

func TestUpdatePoolAddAndDeleteStorage(t *testing.T) {
	poolID := rand.String(10)
	expectedStorage := "local-lvm"
	assert.NoError(t, testClient.CreatePool(poolID, ""))
	assert.NoError(t, testClient.UpdatePool(poolID, nil, &[]string{expectedStorage}, nil, false))
	actualPool, _ := testClient.RetrievePool(poolID)
	actualMembers := actualPool.Members[0].(map[string]any)
	assert.Equal(t, actualMembers["storage"], expectedStorage)

	assert.NoError(t, testClient.UpdatePool(poolID, nil, &[]string{expectedStorage}, nil, true))
	actualPool, _ = testClient.RetrievePool(poolID)
	assert.Len(t, actualPool.Members, 0)
}

func TestUpdatePoolAddAndDeleteVMs(t *testing.T) {
	// TODO: Needs implementing adding VMs/CTs
}
