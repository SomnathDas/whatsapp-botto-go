package main

import (
	"context"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/taigrr/elevenlabs/client"
	"github.com/taigrr/elevenlabs/client/types"
)

func TextToSpeechElevenLabs(text string, fileId string) ([]byte, error) {
	ctx := context.Background()

	// load api key into the client
	client := client.New(goDotEnvVariable("ELEVEN_LABS_TTS_API_KEY"))

	audioBytes, err := client.TTS(ctx, text, goDotEnvVariable("ELEVEN_LABS_VOICE_ID"), goDotEnvVariable("ELEVEN_LABS_MODEL_ID"), types.SynthesisOptions{
		Stability: 0.75, SimilarityBoost: 0.75,
	})

	if err != nil {
		color.Red("FATAL ElevenLabs ERROR: ", err)
	}

	writeErr := os.WriteFile(strings.Join([]string{goDotEnvVariable("AUDIO_FOLDER_ABSOLUTE_PATH"), fileId, ".mp3"}, ""), audioBytes, 0644)

	return audioBytes, writeErr
}
