package xtts

import (
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

type cloneResponse struct {
	*Voice `json:"voice"`
}

func newCloneResponse() *cloneResponse {
	return &cloneResponse{
		Voice: &Voice{},
	}
}

type Voice struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	Favorite     bool   `json:"favorite"`
	SamplesCount int    `json:"samples_count"`
}

func (v *Voice) String() string {
	jsonString, _ := sonic.MarshalString(v)
	return jsonString
}

func (c *Client) CreateVoiceFromFiles(name string, files ...string) (*Voice, error) {
	result := newCloneResponse()
	req := req.C().R().
		SetHeaders(map[string]string{
			"accept":        "application/json",
			"content-type":  "multipart/form-data",
			"authorization": c.authorization,
		}).
		SetSuccessResult(result).
		SetFormData(map[string]string{
			"name": name,
		})

	for _, file := range files {
		req.SetFile("files", file)
	}

	res, err := req.Post("https://app.coqui.ai/api/v2/voices/xtts")
	if err != nil {
		return nil, err
	}
	if res.IsErrorState() {
		return nil, newRequestError(res.StatusCode)
	}
	return result.Voice, nil
}

func (c *Client) CreateVoiceFromUrls(name string, urls ...string) (*Voice, error) {
	result := newCloneResponse()
	payload := map[string]any{
		"name": name,
		"urls": urls,
	}
	res, err := req.C().R().
		SetHeaders(map[string]string{
			"accept":        "application/json",
			"content-type":  "application/json",
			"authorization": c.authorization,
		}).
		SetSuccessResult(result).
		SetBodyJsonMarshal(payload).Post("https://app.coqui.ai/api/v2/voices/xtts/clone-from-url")
	if err != nil {
		return nil, err
	}
	if res.IsErrorState() {
		return nil, newRequestError(res.StatusCode)
	}
	return result.Voice, nil
}

func (c *Client) CreateVoiceFromTextPrompt(name, prompt string) (*Voice, error) {
	result := newCloneResponse()
	payload := map[string]any{
		"name":   name,
		"prompt": prompt,
	}
	res, err := req.C().R().
		SetHeaders(map[string]string{
			"accept":        "application/json",
			"content-type":  "application/json",
			"authorization": c.authorization,
		}).
		SetSuccessResult(result).
		SetBodyJsonMarshal(payload).Post("https://app.coqui.ai/api/v2/voices/xtts/clone-from-url")
	if err != nil {
		return nil, err
	}
	if res.IsErrorState() {
		return nil, newRequestError(res.StatusCode)
	}
	return result.Voice, nil
}
