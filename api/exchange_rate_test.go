package api

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewExchangeRateService(t *testing.T) {
	srv := NewExchangeRateService()
	assert.Equal(t, false, srv.IsProdEnvironment)
}

func TestSetIsProdEnvironmentInExchangeRateService(t *testing.T) {
	srv := dummyExchangeRateService()
	srv.SetIsProdEnvironment(true)
	assert.Equal(t, true, srv.IsProdEnvironment)
}

func TestSetSymbolInExchangeRateService(t *testing.T) {
	srv := dummyExchangeRateService()
	var symbol = "USDT_ERC20"
	srv.SetSymbol(symbol)
	assert.Equal(t, symbol, srv.Request.Symbol)
}

func TestSetTimeStampInExchangeRateService(t *testing.T) {
	srv := dummyExchangeRateService()
	var timestamp = strconv.Itoa(int(time.Now().Unix()))
	srv.SetTimeStamp(timestamp)
	assert.Equal(t, timestamp, srv.Request.TimeStamp)
}

func TestExecuteInExchangeRateService(t *testing.T) {
	srv := dummyExchangeRateService()
	srv.SetIsProdEnvironment(false)
	srv.SetSymbol("USDT_ERC20")
	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, merchantId & signKey is Fake")
	}

	if reflect.TypeOf(rst).String() != "*api.ExchangeRateResponse" {
		t.Fatalf("ExchangeRateService.Execute is expected to return *api.ExchangeRateResponse")
	}
}

func dummyExchangeRateService() *ExchangeRateService {
	return NewExchangeRateService()
}
