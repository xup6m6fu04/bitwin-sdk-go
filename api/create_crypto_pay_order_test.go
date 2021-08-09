package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewCreateCryptoPayOrderService(t *testing.T) {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	srv := NewCreateCryptoPayOrderService(merchantID, signKey)
	assert.Equal(t, merchantID, srv.Request.MerchantID)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantUserIDInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var merchantUserID = "dummyMerchantUserID"
	srv.SetMerchantUserID(merchantUserID)
	assert.Equal(t, merchantUserID, srv.Request.MerchantUserID)
}

func TestSetMerchantOrderIDInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var merchantOrderID = "dummyMerchantOrderID"
	srv.SetMerchantOrderID(merchantOrderID)
	assert.Equal(t, merchantOrderID, srv.Request.MerchantOrderID)
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

func TestSetCallBackURLInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var callBackURL = "dummyCallBackURL"
	srv.SetCallBackURL(callBackURL)
	assert.Equal(t, callBackURL, srv.Request.CallBackURL)
}

func TestSetTimeStampInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInCreateCryptoPayOrderService(t *testing.T) {
	srv := dummyCreateCryptoPayOrderService()
	srv.SetMerchantUserID("YOZERO_USER_000001")
	srv.SetMerchantOrderID("YOZERO_ORDER_TEST_000001")
	srv.SetOrderDescription("YOZERO_DESC_TEST_000001")
	srv.SetAmount("70000000000")
	srv.SetMerchantRMB("4500.35")
	srv.SetSymbol("USDT_ERC20")
	srv.SetCallBackURL("https://attendance.hearts.tw/api/bitwin-payment-callback")
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
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	return NewCreateCryptoPayOrderService(merchantID, signKey)
}
