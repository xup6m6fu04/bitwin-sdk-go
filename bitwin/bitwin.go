package bitwin

import "github.com/xup6m6fu04/bitwin-sdk-go/v3/api"

// Client for BITWIN services
type Client struct {
	MerchantID string
	SignKey    string

	// Services
	CreateCryptoPayOrder  *api.CreateCryptoPayOrderService
	QueryCryptoPayOrder   *api.QueryCryptoPayOrderService
	MerchantWithdraw      *api.MerchantWithdrawService
	QueryMerchantWithdraw *api.QueryMerchantWithdrawService
	ExchangeRate          *api.ExchangeRateService
	BuildRelationUser     *api.BuildRelationUserService
}

// New creates client
func New(merchantID, signKey string) *Client {
	c := &Client{
		MerchantID: merchantID,
		SignKey:    signKey,
	}

	c.CreateCryptoPayOrder = api.NewCreateCryptoPayOrderService(merchantID, signKey)
	c.QueryCryptoPayOrder = api.NewQueryCryptoPayOrderService(merchantID, signKey)
	c.MerchantWithdraw = api.NewMerchantWithdrawService(merchantID, signKey)
	c.QueryMerchantWithdraw = api.NewQueryMerchantWithdrawService(merchantID, signKey)
	c.ExchangeRate = api.NewExchangeRateService()
	c.BuildRelationUser = api.NewBuildRelationUserService(merchantID, signKey)

	return c
}
