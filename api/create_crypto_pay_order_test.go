package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewCreateCryptoPayOrderService(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	srv := NewCreateCryptoPayOrderService(merchantId, signKey)
	assert.Equal(t, merchantId, srv.Request.MerchantId)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantUserIdInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var merchantUserId = "dummyMerchantUserId"
	srv.SetMerchantUserId(merchantUserId)
	assert.Equal(t, merchantUserId, srv.Request.MerchantUserId)
}

func TestSetMerchantOrderIdInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var merchantOrderId = "dummyMerchantOrderId"
	srv.SetMerchantOrderId(merchantOrderId)
	assert.Equal(t, merchantOrderId, srv.Request.MerchantOrderId)
}

func TestSetOrderDescriptionInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var orderDescription = "dummyOrderDescription"
	srv.SetOrderDescription(orderDescription)
	assert.Equal(t, orderDescription, srv.Request.OrderDescription)
}

func TestSetAmountInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var orderAmount = "198964000000"
	srv.SetAmount(orderAmount)
	assert.Equal(t, orderAmount, srv.Request.Amount)
}

func TestSetMerchantRMBInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var merchantRMB = "1989.64"
	srv.SetMerchantRMB(merchantRMB)
	assert.Equal(t, merchantRMB, srv.Request.MerchantRMB)
}

func TestSetSymbolInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var symbol = "USDT_ERC20"
	srv.SetSymbol(symbol)
	assert.Equal(t, symbol, srv.Request.Symbol)
}

func TestSetCallBackUrlInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var callBackUrl = "dummyCallBackUrl"
	srv.SetCallBackUrl(callBackUrl)
	assert.Equal(t, callBackUrl, srv.Request.CallBackUrl)
}

func TestSetTimeStampInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	srv.SetMerchantUserId("YOZERO_USER_000001")
	srv.SetMerchantOrderId("YOZERO_ORDER_TEST_000001")
	srv.SetOrderDescription("YOZERO_DESC_TEST_000001")
	srv.SetAmount("70000000000")
	srv.SetMerchantRMB("4500.35")
	srv.SetSymbol("USDT_ERC20")
	srv.SetCallBackUrl("https://attendance.hearts.tw/api/bitwin-payment-callback")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.CreateCryptoPayOrderResponse" {
		t.Fatalf("CreateCryptoPayOrderService.Execute is expected to return *api.CreateCryptoPayOrderResponse")
	}
}

func dummyCreateCryptoPayOrderService() *CreateCryptoPayOrderService {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	return NewCreateCryptoPayOrderService(merchantId, signKey)
}
