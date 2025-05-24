package antibotaio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SolveSyncEndpoint   = "https://api.antibotaio.dev/solve/nudata/sync"
	SolveWidgetEndpoint = "https://api.antibotaio.dev/solve/nudata/widget/%v"
)

type SyncInput struct {
	Href      string                 `json:"href"`
	Website   string                 `json:"website"`
	Body      string                 `json:"body"`
	UserAgent string                 `json:"user_agent"`
	Language  string                 `json:"language"`
	Timezone  string                 `json:"timezone"`
	Args      map[string]interface{} `json:"args"`
}

type NuDataWidgetTask struct {
	Body string `json:"body"`
}

type SolveSyncResponse struct {
	ID       string `json:"id"`
	Solution string `json:"solution"`
}

type SolveWidgetResponse struct {
	Solution string `json:"solution"`
}

func (s *Session) SolveSync(input *SyncInput) (*SolveSyncResponse, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", SolveSyncEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", s.APIKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to solve sync: %v", resp.Status)
	}

	defer resp.Body.Close()

	var response SolveSyncResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (s *Session) SolveWidget(input *NuDataWidgetTask, taskID string) (*SolveWidgetResponse, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(SolveWidgetEndpoint, taskID), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", s.APIKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response SolveWidgetResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
