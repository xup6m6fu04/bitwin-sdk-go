package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewQueryCryptoPayOrderService(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	srv := NewQueryCryptoPayOrderService(merchantId, signKey)
	assert.Equal(t, merchantId, srv.Request.MerchantId)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantOrderIdInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	var dummyMerchantOrderId = "dummyMerchantOrderId"
	srv.SetMerchantOrderId("dummyMerchantOrderId")
	assert.Equal(t, dummyMerchantOrderId, srv.Request.MerchantOrderId)
}

func TestSetOrderIdInQueryCryptoPayOrderService(t *testing.T) {
	srv := dummyQueryCryptoPayOrderService()
	var dummyOrderId = "dummyOrderId"
	srv.SetOrderId("dummyOrderId")
	assert.Equal(t, dummyOrderId, srv.Request.OrderId)
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
	srv.SetMerchantOrderId("YOZERO_ORDER_000002")
	srv.SetOrderId("53110728501176672")
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
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	return NewQueryCryptoPayOrderService(merchantId, signKey)
}