package eleven

import (
	"fmt"
	"net/http"
	"os"
)

type Eleven struct {
	baseURL string
	client  *http.Client
	secret  string
}

type StreamingLatency int

const (
	// no latency optimizations
	StreamingLatencyDefault StreamingLatency = 0

	// 50% latency optimizations
	StreamingLatencyNormal StreamingLatency = 1

	// 75% latency optimizations
	StreamingLatencyStrong StreamingLatency = 2

	// max latency optimizations, but also with text normalizer turned off for even more latency savings (best latency, but can mispronounce eg numbers and dates).
	StreamingLatencyMax StreamingLatency = 3
)

type VoiceSettings struct {
	SimilarityBoost int `json:"similarity_boost"`
	Stability       int `json:"stability"`
}

// You can find reference to the settings here: https://docs.elevenlabs.io/api-reference/text-to-speech#query
type GenerateArgs struct {
	// Text to be converted to speech
	Text string

	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices or call the .ListVoices() method to list all the available voices.
	VoiceID string

	// Identifier of the model that will be used - default is "eleven_monolingual_v1"
	ModelID string

	// You can turn on latency optimizations at some cost of quality. The best possible final latency varies by model
	OptimiseStreamingLatency StreamingLatency

	// Voice settings overriding stored settings for the given voice. They are applied only on the given TTS request.
	VoiceSettings VoiceSettings
}

func New(args ...string) *Eleven {
	key := os.Getenv("ELEVEN_API_KEY")
	if len(args) > 0 {
		key = args[0]
	}
	if key == "" {
		fmt.Println("API key not set, your requests will be limited")
	}
	return &Eleven{
		baseURL: "https://api.elevenlabs.io/v1",
		client:  http.DefaultClient,
		secret:  key,
	}
}

// SetSecret sets the secret to be used for all requests after the point it is called.
func (e *Eleven) SetSecret(secret string) {
	e.secret = secret
}

// Generate generates audio from the given text.
func (e *Eleven) Generate(args *GenerateArgs) (string, error) {
	var path string

	return path, nil
}
