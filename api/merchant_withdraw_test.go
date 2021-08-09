package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewMerchantWithdrawService(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	srv := NewMerchantWithdrawService(merchantId, signKey)
	assert.Equal(t, merchantId, srv.Request.MerchantId)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantWithdrawIdInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var merchantWithdrawId = "dummyMerchantWithdrawId"
	srv.SetMerchantWithdrawId(merchantWithdrawId)
	assert.Equal(t, merchantWithdrawId, srv.Request.MerchantWithdrawId)
}

func TestSetUserWalletInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var userWallet = "0x875EDa094F03Ed4c93adb3dbb77913F860dC888f"
	srv.SetUserWallet(userWallet)
	assert.Equal(t, userWallet, srv.Request.UserWallet)
}

func TestSetAmountInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var orderAmount = "198964000000"
	srv.SetAmount(orderAmount)
	assert.Equal(t, orderAmount, srv.Request.Amount)
}

func TestSetMerchantRMBInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var merchantRMB = "1989.64"
	srv.SetMerchantRMB(merchantRMB)
	assert.Equal(t, merchantRMB, srv.Request.MerchantRMB)
}

func TestSetSymbolInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var symbol = "USDT_ERC20"
	srv.SetSymbol(symbol)
	assert.Equal(t, symbol, srv.Request.Symbol)
}

func TestSetCallBackUrlInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var callBackUrl = "dummyCallBackUrl"
	srv.SetCallBackUrl(callBackUrl)
	assert.Equal(t, callBackUrl, srv.Request.CallBackUrl)
}

func TestSetTimeStampInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	srv.SetMerchantUserId("YOZERO_USER_000001")
	srv.SetMerchantWithdrawId("YOZERO_WITHDRAW_000003")
	srv.SetUserWallet("0x875EDa094F03Ed4c93adb3dbb77913F860dC888f")
	srv.SetAmount("1000000000")
	srv.SetMerchantRMB("66.28")
	srv.SetSymbol("USDT_ERC20")
	srv.SetCallBackUrl("https://attendance.hearts.tw/api/bitwin-withdraw-callback")
	srv.SetTimeStamp(strconv.Itoa(int(time.Now().Unix())))
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.MerchantWithdrawResponse" {
		t.Fatalf("MerchantWithdrawService.Execute is expected to return *api.MerchantWithdrawResponse")
	}
}

func dummyMerchantWithdrawService() *MerchantWithdrawService {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	return NewMerchantWithdrawService(merchantId, signKey)
}
