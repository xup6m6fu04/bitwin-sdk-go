package bitwin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	merchantID := "DummyMerchantId"
	signKey := "DummySighKey"
	accessKey := "DummyAccessKey"
	c := New(merchantID, signKey, accessKey)
	assert.Equal(t, signKey, c.CreateCryptoPayOrder.SignKey)
	assert.Equal(t, signKey, c.QueryCryptoPayOrder.SignKey)
	assert.Equal(t, signKey, c.MerchantWithdraw.SignKey)
	assert.Equal(t, signKey, c.QueryMerchantWithdraw.SignKey)
	assert.Equal(t, false, c.ExchangeRate.IsProdEnvironment)
	assert.Equal(t, signKey, c.BuildRelationUser.SignKey)
	assert.Equal(t, merchantID, c.CreateCryptoPayOrder.Request.MerchantID)
	assert.Equal(t, merchantID, c.QueryCryptoPayOrder.Request.MerchantID)
	assert.Equal(t, merchantID, c.MerchantWithdraw.Request.MerchantID)
	assert.Equal(t, merchantID, c.QueryMerchantWithdraw.Request.MerchantID)
	assert.Equal(t, merchantID, c.BuildRelationUser.Request.MerchantID)
	assert.Equal(t, accessKey, c.CreateCryptoPayOrder.AccessKey)
	assert.Equal(t, accessKey, c.QueryCryptoPayOrder.AccessKey)
	assert.Equal(t, accessKey, c.MerchantWithdraw.AccessKey)
	assert.Equal(t, accessKey, c.QueryMerchantWithdraw.AccessKey)
	assert.Equal(t, accessKey, c.BuildRelationUser.AccessKey)
}
