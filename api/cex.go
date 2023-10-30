package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	dt "frontendmasters.com/go/crypto/datatypes"
)

const ApiUrl = "https://cex.io/api/ticker/%s/USD"

// when dealing with structuers, use a pointer for the return type, so you can return nil
// that way you don't get any complaints
func GetRate(currency string) (*dt.Rate, error)  {

	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required; %d received", len(currency))
	}

	capCurrency := strings.ToUpper(currency)
	response, err :=  http.Get(fmt.Sprintf(ApiUrl, capCurrency))

	// this is a network error, like the domain is down or the url is invalid
	if err != nil {
		return nil, err
	}

	var marshalResponse CexResponse

	if response.StatusCode == http.StatusOK {
		// readAll reads the incoming response as chunks of bytes, so we read the response as it comes in
		bodyBytes, bodyErr := io.ReadAll(response.Body)

		// a possible error here is that the server closed the connections
		if bodyErr != nil {
			return nil, bodyErr
		}

		// var cryptoRate dt.Rate
		jsonParseErr := json.Unmarshal(bodyBytes, &marshalResponse)

		if jsonParseErr != nil {
			return nil, jsonParseErr
		}

	} else {
		return nil, fmt.Errorf("Status code received %v", response.StatusCode)
	}

	rate := dt.Rate {Currency: currency, Price: marshalResponse.Bid}

	return &rate, nil
}