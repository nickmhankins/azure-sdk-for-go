// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// HTTPPoller provides polling facilities until the operation completes
type HTTPPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*http.Response, error)
	ResumeToken() (string, error)
}

type httpPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, p.pipeline, nil)
}

// ResumeToken generates the string token that can be used with the ResumeHTTPPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, frequency, p.pipeline, nil)
}

// PrivateEndpointConnectionPoller provides polling facilities until the operation completes
type PrivateEndpointConnectionPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionResponse, error)
	ResumeToken() (string, error)
}

type privateEndpointConnectionPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *privateEndpointConnectionPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *privateEndpointConnectionPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *privateEndpointConnectionPoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionResponse, error) {
	respType := PrivateEndpointConnectionResponse{PrivateEndpointConnection: &PrivateEndpointConnection{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumePrivateEndpointConnectionPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *privateEndpointConnectionPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (PrivateEndpointConnectionResponse, error) {
	respType := PrivateEndpointConnectionResponse{PrivateEndpointConnection: &PrivateEndpointConnection{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// VaultPoller provides polling facilities until the operation completes
type VaultPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (VaultResponse, error)
	ResumeToken() (string, error)
}

type vaultPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *vaultPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *vaultPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *vaultPoller) FinalResponse(ctx context.Context) (VaultResponse, error) {
	respType := VaultResponse{Vault: &Vault{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.Vault)
	if err != nil {
		return VaultResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumeVaultPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *vaultPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *vaultPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (VaultResponse, error) {
	respType := VaultResponse{Vault: &Vault{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.Vault)
	if err != nil {
		return VaultResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
