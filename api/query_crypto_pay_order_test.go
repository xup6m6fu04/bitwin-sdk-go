package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewQueryCryptoPayOrderService(t *testing.T) {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	accessKey := "DummyAccessKey"
	srv := NewQueryCryptoPayOrderService(merchantID, signKey, accessKey)
	assert.Equal(t, merchantID, srv.Request.MerchantID)
	assert.Equal(t, signKey, srv.SignKey)
	assert.Equal(t, accessKey, srv.AccessKey)
}

func TestSetIsProdEnvironmentInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantOrderIDInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	var dummyMerchantOrderID = "dummyMerchantOrderID"
	srv.SetMerchantOrderID("dummyMerchantOrderID")
	assert.Equal(t, dummyMerchantOrderID, srv.Request.MerchantOrderID)
}

func TestSetOrderIDInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	var dummyOrderID = "dummyOrderID"
	srv.SetOrderID("dummyOrderID")
	assert.Equal(t, dummyOrderID, srv.Request.OrderID)
}

func TestSetTimeStampInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	srv.SetIsProdEnvironment(false)
	srv.SetMerchantOrderID("YOZERO_ORDER_000002")
	srv.SetOrderID("53110728501176672")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.QueryCryptoPayOrderResponse" {
		t.Fatalf("QueryCryptoPayOrderService.Execute is expected to return *api.QueryCryptoPayOrderResponse")
	}
}

func dummyQueryCryptoPayOrderService() *QueryCryptoPayOrderService {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	accessKey := "DummyAccessKey"
	return NewQueryCryptoPayOrderService(merchantID, signKey, accessKey)
}
