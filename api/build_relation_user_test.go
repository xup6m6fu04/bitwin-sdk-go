package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewBuildRelationUserService(t *testing.T) {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	srv := NewBuildRelationUserService(merchantID, signKey)
	assert.Equal(t, merchantID, srv.Request.MerchantID)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantUserIDInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	var merchantUserID = "dummyMerchantUserID"
	srv.SetMerchantUserID(merchantUserID)
	assert.Equal(t, merchantUserID, srv.Request.MerchantUserID)
}

func TestSetCallBackURLInBuildRelationUserService(t *testing.T) {
	srv := dummyBuildRelationUserService()
	var callBackURL = "dummyCallBackURL"
	srv.SetCallBackURL(callBackURL)
	assert.Equal(t, callBackURL, srv.Request.CallBackURL)
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
	srv.SetMerchantUserID("YOZERO_USER_000001")
	srv.SetCallBackURL("https://attendance.hearts.tw/api/bitwin-withdraw-callback")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantID & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.BuildRelationUserResponse" {
		t.Fatalf("BuildRelationUserService.Execute is expected to return *api.BuildRelationUserResponse")
	}
}

func dummyBuildRelationUserService() *BuildRelationUserService {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	return NewBuildRelationUserService(merchantID, signKey)
}
