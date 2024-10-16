package contractevent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/skip-mev/connect/v2/oracle/types"
)

// ContractEventFetcher fetches events from Ethereum contracts using an EventProvider.
type ContractEventFetcher struct {
	eventProvider types.EventProvider // Use the EventProvider to fetch contract events
}

// NewContractEventFetcher initializes a new ContractEventFetcher with the given EventProvider.
func NewContractEventFetcher(provider types.EventProvider) *ContractEventFetcher {
	return &ContractEventFetcher{
		eventProvider: provider,
	}
}

// FetchEvents fetches contract events from the Ethereum network.
func (f *ContractEventFetcher) FetchEvents() ([]byte, error) {
	fromBlock := "0x0"
	toBlock := "latest"

	requestData := types.EthLogRequest{
		JsonRPC: "2.0",
		Method:  "eth_getLogs",
		Params: []interface{}{
			map[string]interface{}{
				"fromBlock": fromBlock,
				"toBlock":   toBlock,
				"address":   f.eventProvider.ContractConfig.Address,
				"topics":    f.eventProvider.ContractConfig.Topics,
			},
		},
		ID: 1,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Try each endpoint in the provider's RPC config
	for _, endpoint := range f.eventProvider.EthereumRPCConfig.Endpoints {
		result, err := f.makeRPCCallWithTimeout(endpoint, requestBody, f.eventProvider.EthereumRPCConfig.Timeout)
		if err == nil {
			return result, nil // Return the result if the call is successful
		}
		fmt.Printf("Failed to fetch from endpoint %s: %v\n", endpoint, err)
	}

	return nil, errors.New("all endpoints failed")
}

// makeRPCCallWithTimeout makes the actual RPC call with a timeout from EthereumRPCConfig
func (f *ContractEventFetcher) makeRPCCallWithTimeout(endpoint string, requestBody []byte, timeout int) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second, // Set timeout from RPC config
	}

	resp, err := client.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return body, nil
}
