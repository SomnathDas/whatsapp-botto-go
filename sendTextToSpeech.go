package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func SendTextToSpeech(WAClient *whatsmeow.Client, v *events.Message, aiResponse string) {
	// Showing [bot... is recording audio....] in whatsapp until audio file is ready
	WAClient.SendChatPresence(v.Info.Chat, types.ChatPresenceComposing, types.ChatPresenceMediaAudio)

	// Converting AI Response to Speech (using ElevenLabs API)
	// TextToSpeechElevenLabs function also returns []byte but for now I just let it store mp3 file directly
	_, writeErr := TextToSpeechElevenLabs(aiResponse, v.Info.ID)
	if writeErr != nil {
		color.Red("TEXT TO SPEECH WRITE ERROR: ", writeErr)
	}

	// converting generated mp3 file to ogg opus file to send to whatsapp servers
	// ConvertAudio() function uses "ffmpeg along with libopus" through exec.Command()
	ConvertAudio(v.Info.ID)

	// Reading the audio file created by TextToSpeech func then converted to opus file by ConvertAudio func
	audioBytes, err := os.ReadFile(strings.Join([]string{goDotEnvVariable("AUDIO_FOLDER_ABSOLUTE_PATH"), v.Info.ID, ".opus"}, ""))

	if err != nil {
		fmt.Println(err)
	}

	// Uploading the audio to the whatsapp server using whatsmeow library method
	audioUploaded, err := WAClient.Upload(context.Background(), audioBytes, whatsmeow.MediaAudio)
	if err != nil {
		err := errors.New("error while uploading media to whatsapp server")
		fmt.Println(err)
	}

	//  Composing the response message to be sent
	response, err := WAClient.SendMessage(context.Background(), v.Info.Chat, &waProto.Message{
		AudioMessage: &waProto.AudioMessage{
			Url:           proto.String(audioUploaded.URL),
			Ptt:           proto.Bool(true),
			DirectPath:    proto.String(audioUploaded.DirectPath),
			Mimetype:      proto.String("audio/ogg; codecs=opus"),
			FileLength:    proto.Uint64(audioUploaded.FileLength),
			FileSha256:    audioUploaded.FileSHA256,
			FileEncSha256: audioUploaded.FileEncSHA256,
			MediaKey:      audioUploaded.MediaKey,
			// Providing ContextInfo to send this message as a "reply"
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      &v.Info.ID,
				QuotedMessage: v.Message.ExtendedTextMessage.ContextInfo.GetQuotedMessage(),
				Participant:   proto.String(v.Info.MessageSource.Sender.String()),
			},
		},
	})

	if err != nil {
		color.Red("\n\nMessage Sending Failure: %v\n\n", err)
	}

	color.Green("\n\nMessage Sent :\n timestamp: %v \n from %v\n\n", response.Timestamp, "bot")

	// Using CleanAudioFolder func to empty the audio folder
	defer CleanAudioFolder()
}
