package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewQueryMerchantWithdrawService(t *testing.T) {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	srv := NewQueryMerchantWithdrawService(merchantID, signKey)
	assert.Equal(t, merchantID, srv.Request.MerchantID)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetWithdrawIDInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var withdrawID = "dummyWithdrawID"
	srv.SetWithdrawID(withdrawID)
	assert.Equal(t, withdrawID, srv.Request.WithdrawID)
}

func TestSetMerchantWithdrawIDInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var merchantWithdrawID = "dummyMerchantWithdrawID"
	srv.SetMerchantWithdrawID(merchantWithdrawID)
	assert.Equal(t, merchantWithdrawID, srv.Request.MerchantWithdrawID)
}

func TestSetTimeStampInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	srv.SetWithdrawID("53116499398961504")
	srv.SetMerchantWithdrawID("YOZERO_WITHDRAW_000003")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantID & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.QueryMerchantWithdrawResponse" {
		t.Fatalf("QueryMerchantWithdrawService.Execute is expected to return *api.QueryMerchantWithdrawResponse")
	}
}

func dummyQueryMerchantWithdrawService() *QueryMerchantWithdrawService {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	return NewQueryMerchantWithdrawService(merchantID, signKey)
}
