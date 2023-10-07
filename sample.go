package xtts

import (
	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
)

type Sample struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
	Language  string `json:"language"`
	AudioUrl  string `json:"audio_url"`
}

func (s *Sample) String() string {
	jsonString, _ := sonic.MarshalString(s)
	return jsonString
}

type SampleConf struct {
	VoiceID  string
	Text     string
	Speed    float64
	Language string
}

func DefaultSampleConf(id string, text string) *SampleConf {
	return &SampleConf{
		VoiceID:  id,
		Text:     text,
		Speed:    1,
		Language: "en",
	}
}

func (c *Client) CreateSample(conf *SampleConf) (*Sample, error) {
	result := new(Sample)
	payload := map[string]any{
		"voice_id": conf.VoiceID,
		"text":     conf.Text,
		"speed":    conf.Speed,
		"language": conf.Language,
	}
	res, err := req.C().R().
		SetHeaders(map[string]string{
			"accept":        "application/json",
			"content-type":  "application/json",
			"authorization": c.authorization,
		}).
		SetSuccessResult(result).
		SetBodyJsonMarshal(payload).Post("https://app.coqui.ai/api/v2/samples/xtts")
	if err != nil {
		return nil, err
	}
	if res.IsErrorState() {
		return nil, newRequestError(res.StatusCode)
	}
	return result, nil
}

// func (c *Client) CreateStreamSample(conf *SampleConf) (*Sample, error) {
// 	result := new(Sample)
// 	payload := map[string]any{
// 		"voice_id": conf.VoiceID,
// 		"text":     conf.Text,
// 		"speed":    conf.Speed,
// 		"language": conf.Language,
// 	}
// 	_, err := req.C().R().
// 		SetHeaders(map[string]string{
// 			"accept":        "application/json",
// 			"content-type":  "application/json",
// 			"authorization": c.authorization,
// 		}).
// 		SetSuccessResult(result).
// 		SetBodyJsonMarshal(payload).Post("https://app.coqui.ai/api/v2/voices/xtts/clone-from-url")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
