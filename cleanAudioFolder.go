package main

import (
	"os"

	"github.com/fatih/color"
)

func CleanAudioFolder() {
	err := os.RemoveAll(goDotEnvVariable("AUDIO_FOLDER_ABSOLUTE_PATH"))
	if err != nil {
		color.Red("\n\nERROR DURING DELETING audio folder contents: \n\n", err)
	}
	mkdirErr := os.Mkdir(goDotEnvVariable("AUDIO_FOLDER_ABSOLUTE_PATH"), 0777)
	if mkdirErr != nil {
		color.Red("\n\nERROR DURING CREATING audio folder: \n\n", err)
	}
	color.Green("\n\nSuccessfully cleaned audio folder.\n\n")
}
