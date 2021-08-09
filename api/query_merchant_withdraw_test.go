package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewQueryMerchantWithdrawService(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	srv := NewQueryMerchantWithdrawService(merchantId, signKey)
	assert.Equal(t, merchantId, srv.Request.MerchantId)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetWithdrawIdInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var withdrawId = "dummyWithdrawId"
	srv.SetWithdrawId(withdrawId)
	assert.Equal(t, withdrawId, srv.Request.WithdrawId)
}

func TestSetMerchantWithdrawIdInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var merchantWithdrawId = "dummyMerchantWithdrawId"
	srv.SetMerchantWithdrawId(merchantWithdrawId)
	assert.Equal(t, merchantWithdrawId, srv.Request.MerchantWithdrawId)
}

func TestSetTimeStampInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInQueryMerchantWithdrawService(t *testing.T) {
	srv := dummyQueryMerchantWithdrawService()
	srv.SetWithdrawId("53116499398961504")
	srv.SetMerchantWithdrawId("YOZERO_WITHDRAW_000003")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.QueryMerchantWithdrawResponse" {
		t.Fatalf("QueryMerchantWithdrawService.Execute is expected to return *api.QueryMerchantWithdrawResponse")
	}
}

func dummyQueryMerchantWithdrawService() *QueryMerchantWithdrawService {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	return NewQueryMerchantWithdrawService(merchantId, signKey)
}
