package eleven

import (
	"fmt"
)

type FineTuningState string

// Fine-tuning states
const (
	FineTuningStateNotStarted FineTuningState = "not_started"
	FineTuningStateInProgress FineTuningState = "is_fine_tuning"
	FineTuningStateFinished   FineTuningState = "fine_tuned"
)

type RecordingResponseModel struct {
	MimeType       string `json:"mime_type"`
	RecordingID    string `json:"recording_id"`
	SizeBytes      int    `json:"size_bytes"`
	Transcription  string `json:"transcription"`
	UploadDateUnix int    `json:"upload_date_unix"`
}

type VerificationAttemptsResponseModel struct {
	Accepted            bool                   `json:"accepted"`
	DateUnix            int                    `json:"date_unix"`
	LevenshteinDistance int                    `json:"levenshtein_distance"`
	Recording           RecordingResponseModel `json:"recording"`
	Similarity          int                    `json:"similarity"`
	Text                string                 `json:"text"`
}

type FineTuningResponseModel struct {
	FineTuningRequested       bool                                `json:"fine_tuning_requested"`
	FineTuningState           FineTuningState                     `json:"finetuning_state"`
	IsAllowedToFineTune       bool                                `json:"is_allowed_to_fine_tune"`
	ModelID                   string                              `json:"model_id"`
	SliceIDs                  []string                            `json:"slice_ids"`
	VerificationAttempts      []VerificationAttemptsResponseModel `json:"verification_attempts"`
	VerificationAttemptsCount int                                 `json:"verification_attempts_count"`
	VerificationFailures      []string                            `json:"verification_failures"`
}

type SampleResponseModel struct {
	FileName  string `json:"file_name"`
	Hash      string `json:"hash"`
	MimeType  string `json:"mime_type"`
	SampleID  string `json:"sample_id"`
	SizeBytes int    `json:"size_bytes"`
}

type VoiceResponse struct {
	Name              string                  `json:"name"`
	AvailableForTiers []string                `json:"available_for_tiers"`
	Category          string                  `json:"category"`
	Description       string                  `json:"description"`
	FineTuning        FineTuningResponseModel `json:"fine_tuning"`
	Labels            map[string]any          `json:"labels"`
	PreviewURL        string                  `json:"preview_url"`
	Samples           []SampleResponseModel   `json:"samples"`
	Settings          VoiceSettings           `json:"settings"`
	VoiceID           string                  `json:"voice_id"`
}

func (e *Eleven) ListVoices() ([]VoiceResponse, error) {
	var response struct {
		Voices []VoiceResponse `json:"voices"`
	}

	resp, err := e.get(Request{
		Path: "/voices",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list voices: %w", err)
	}

	err = decodeResponse(resp, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Voices, nil
}
