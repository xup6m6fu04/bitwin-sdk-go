package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewMerchantWithdrawService(t *testing.T) {
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	srv := NewMerchantWithdrawService(merchantID, signKey)
	assert.Equal(t, merchantID, srv.Request.MerchantID)
	assert.Equal(t, signKey, srv.SignKey)
}

func TestSetIsProdEnvironmentInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetMerchantWithdrawIDInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var merchantWithdrawID = "dummyMerchantWithdrawID"
	srv.SetMerchantWithdrawID(merchantWithdrawID)
	assert.Equal(t, merchantWithdrawID, srv.Request.MerchantWithdrawID)
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

func TestSetCallBackURLInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var callBackURL = "dummyCallBackURL"
	srv.SetCallBackURL(callBackURL)
	assert.Equal(t, callBackURL, srv.Request.CallBackURL)
}

func TestSetTimeStampInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInMerchantWithdrawService(t *testing.T) {
	srv := dummyMerchantWithdrawService()
	srv.SetMerchantUserID("YOZERO_USER_000001")
	srv.SetMerchantWithdrawID("YOZERO_WITHDRAW_000003")
	srv.SetUserWallet("0x875EDa094F03Ed4c93adb3dbb77913F860dC888f")
	srv.SetAmount("1000000000")
	srv.SetMerchantRMB("66.28")
	srv.SetSymbol("USDT_ERC20")
	srv.SetCallBackURL("https://attendance.hearts.tw/api/bitwin-withdraw-callback")
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
	merchantID := "DummyMerchantID"
	signKey := "DummySighKey"
	return NewMerchantWithdrawService(merchantID, signKey)
}
