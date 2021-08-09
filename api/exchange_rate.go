package api

import "encoding/json"

type ExchangeRateService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           ExchangeRateRequest
}

type ExchangeRateRequest struct {
	Symbol    string `json:"Symbol"`
	TimeStamp string `json:"TimeStamp"`
}

type ExchangeRateResponse struct {
	RMBRate       string `json:"RMBRate,omitempty"`
	RMBBuyRate    string `json:"RMBBuyRate,omitempty"`
	ReturnCode    string `json:"ReturnCode"`
	ReturnMessage string `json:"ReturnMessage"`
}

func NewExchangeRateService() *ExchangeRateService {
	return &ExchangeRateService{
		IsProdEnvironment: false,
		Request: ExchangeRateRequest{
			Symbol:    "",
			TimeStamp: "",
		},
	}
}

func (srv *ExchangeRateService) SetIsProdEnvironment(IsProdEnvironment bool) *ExchangeRateService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

func (srv *ExchangeRateService) SetSymbol(symbol string) *ExchangeRateService {
	srv.Request.Symbol = symbol
	return srv
}

func (srv *ExchangeRateService) SetTimeStamp(timeStamp string) *ExchangeRateService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

func (srv *ExchangeRateService) Execute() (*ExchangeRateResponse, error) {

	url := "ExchangeRate"

	// Send Request
	body, err := callServer(srv.Request, srv.IsProdEnvironment, url)
	var resp ExchangeRateResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
