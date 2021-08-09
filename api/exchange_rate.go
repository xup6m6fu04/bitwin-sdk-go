package api

import "encoding/json"

// ExchangeRateService type
type ExchangeRateService struct {
	IsProdEnvironment bool
	SignKey           string
	Request           ExchangeRateRequest
}

// ExchangeRateRequest type
type ExchangeRateRequest struct {
	Symbol    string `json:"Symbol"`
	TimeStamp string `json:"TimeStamp"`
}

// ExchangeRateResponse type
type ExchangeRateResponse struct {
	RMBRate       string `json:"RMBRate,omitempty"`
	RMBBuyRate    string `json:"RMBBuyRate,omitempty"`
	ReturnCode    string `json:"ReturnCode"`
	ReturnMessage string `json:"ReturnMessage"`
}

// NewExchangeRateService returns a new service
func NewExchangeRateService() *ExchangeRateService {
	return &ExchangeRateService{
		IsProdEnvironment: false,
		Request: ExchangeRateRequest{
			Symbol:    "",
			TimeStamp: "",
		},
	}
}

// SetIsProdEnvironment sets the environment (test or prod)
func (srv *ExchangeRateService) SetIsProdEnvironment(IsProdEnvironment bool) *ExchangeRateService {
	srv.IsProdEnvironment = IsProdEnvironment
	return srv
}

// SetSymbol sets Symbol
func (srv *ExchangeRateService) SetSymbol(symbol string) *ExchangeRateService {
	srv.Request.Symbol = symbol
	return srv
}

// SetTimeStamp sets Timestamp
func (srv *ExchangeRateService) SetTimeStamp(timeStamp string) *ExchangeRateService {
	srv.Request.TimeStamp = timeStamp
	return srv
}

// Execute build and send request(*ExchangeRateService) return *ExchangeRateResponse
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
