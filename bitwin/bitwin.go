package bitwin

import "github.com/xup6m6fu04/bitwin-sdk-go/api"

type Client struct {
	MerchantId string
	SignKey    string

	// Services
	CreateCryptoPayOrder  *api.CreateCryptoPayOrderService
	QueryCryptoPayOrder   *api.QueryCryptoPayOrderService
	MerchantWithdraw      *api.MerchantWithdrawService
	QueryMerchantWithdraw *api.QueryMerchantWithdrawService
	ExchangeRate          *api.ExchangeRateService
	BuildRelationUser     *api.BuildRelationUserService
}

func New(merchantId, signKey string) *Client {
	c := &Client{
		MerchantId: merchantId,
		SignKey:    signKey,
	}

	c.CreateCryptoPayOrder = api.NewCreateCryptoPayOrderService(merchantId, signKey)
	c.QueryCryptoPayOrder = api.NewQueryCryptoPayOrderService(merchantId, signKey)
	c.MerchantWithdraw = api.NewMerchantWithdrawService(merchantId, signKey)
	c.QueryMerchantWithdraw = api.NewQueryMerchantWithdrawService(merchantId, signKey)
	c.ExchangeRate = api.NewExchangeRateService()
	c.BuildRelationUser = api.NewBuildRelationUserService(merchantId, signKey)

	return c
}
