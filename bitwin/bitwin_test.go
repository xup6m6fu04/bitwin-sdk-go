package bitwin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	merchantId := "DummyMerchantId"
	signKey := "DummySighKey"
	c := New(merchantId, signKey)
	assert.Equal(t, signKey, c.CreateCryptoPayOrder.SignKey)
	assert.Equal(t, signKey, c.QueryCryptoPayOrder.SignKey)
	assert.Equal(t, signKey, c.MerchantWithdraw.SignKey)
	assert.Equal(t, signKey, c.QueryMerchantWithdraw.SignKey)
	assert.Equal(t, false, c.ExchangeRate.IsProdEnvironment)
	assert.Equal(t, signKey, c.BuildRelationUser.SignKey)
	assert.Equal(t, merchantId, c.CreateCryptoPayOrder.Request.MerchantId)
	assert.Equal(t, merchantId, c.QueryCryptoPayOrder.Request.MerchantId)
	assert.Equal(t, merchantId, c.MerchantWithdraw.Request.MerchantId)
	assert.Equal(t, merchantId, c.QueryMerchantWithdraw.Request.MerchantId)
	assert.Equal(t, merchantId, c.BuildRelationUser.Request.MerchantId)
}
