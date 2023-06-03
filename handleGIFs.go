package main

import (
	"github.com/fatih/color"
	"go.mau.fi/whatsmeow/types/events"
)

func HandleGIFs(v *events.Message) {
	color.Blue("\n\nGIF Recieved :\n %v \n from %v\n\n", v.Message.StickerMessage.GetMimetype(), v.Info.Sender)
}
