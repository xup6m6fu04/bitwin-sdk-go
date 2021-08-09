package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewBuildRelationUserService(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	srv := NewBuildRelationUserService(merchantId, signKey)
	assert.Equal(t, merchantId, srv.Request.MerchantId)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantUserIdInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	var merchantUserId = "dummyMerchantUserId"
	srv.SetMerchantUserId(merchantUserId)
	assert.Equal(t, merchantUserId, srv.Request.MerchantUserId)
}

func TestSetCallBackUrlInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	var callBackUrl = "dummyCallBackUrl"
	srv.SetCallBackUrl(callBackUrl)
	assert.Equal(t, callBackUrl, srv.Request.CallBackUrl)
}

func TestSetTimeStampInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	srv.SetIsProdEnvironment(false)
	srv.SetMerchantUserId("YOZERO_USER_000001")
	srv.SetCallBackUrl("https://attendance.hearts.tw/api/bitwin-withdraw-callback")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.BuildRelationUserResponse" {
		t.Fatalf("BuildRelationUserService.Execute is expected to return *api.BuildRelationUserResponse")
	}
}

func dummyBuildRelationUserService() *BuildRelationUserService {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	return NewBuildRelationUserService(merchantId, signKey)
}


