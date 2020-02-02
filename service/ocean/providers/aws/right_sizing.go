package aws

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

//Suggestion - definition of suggestion ouptut of Spot API
type Suggestion struct {
	DeploymentName  *string `json:"deploymentName,omitempty"`
	Namespace       *string `json:"namespace,omitempty"`
	SuggestedCPU    *int    `json:"suggestedCPU,omitempty"`
	RequestedCPU    *int    `json:"requestedCPU,omitempty"`
	SuggestedMemory *int    `json:"suggestedMemory,omitempty"`
	RequestedMemory *int    `json:"requestedMemory,omitempty"`
}

//ReadRightSizingInput - Input struct required for getting Spot Right
//Sizing suggestions for an Ocean cluster
type ReadRightSizingInput struct {
	OceanID *string `json:"oceanId,omitempty"`
}

//ReadRightSizingOutput - output struct of suggestion array as Right Sizing
//API response with array of suggestions per Namespace & Deploymnet
type ReadRightSizingOutput struct {
	Suggestions []*Suggestion `json:"suggestions,omitempty"`
}

func rightSizingFromJSON(in []byte) (*ReadRightSizingOutput, error) {
	b := new(ReadRightSizingOutput)

	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
func rightSizingsFromJSON(in []byte) (*ReadRightSizingOutput, error) {
	var rw client.RsResponse
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	a, _ := rightSizingFromJSON(rw.Response.Items)
	return a, nil
}

func rightSizingFromHTTPResponse(resp *http.Response) (*ReadRightSizingOutput, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rightSizingsFromJSON(body)
}

//ReadRightSizing - get all right sizing suggestions for a single Ocean cluster
func (s *ServiceOp) ReadRightSizing(ctx context.Context, input *ReadRightSizingInput) (*ReadRightSizingOutput, error) {
	path, err := uritemplates.Expand("/ocean/aws/k8s/cluster/{oceanId}/rightSizing/suggestion", uritemplates.Values{
		"oceanId": spotinst.StringValue(input.OceanID),
	})
	if err != nil {
		return nil, err
	}
	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := rightSizingFromHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	return gs, nil
}