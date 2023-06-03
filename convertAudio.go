package main

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func ConvertAudio(fileId string) {
	_, err := exec.Command("ffmpeg", "-i", strings.Join([]string{"./audio/", fileId, ".mp3"}, ""), "-c:a", "libopus", strings.Join([]string{"./audio/", fileId, ".opus"}, "")).Output()
	if err != nil {
		color.Red("Audio Conversion Error:", err)
	}
	color.Green("Audio Conversion Success for fileId:", fileId)

}
