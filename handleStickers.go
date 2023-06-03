package main

import (
	"github.com/fatih/color"
	"go.mau.fi/whatsmeow/types/events"
)

func HandleStickers(v *events.Message) {
	color.Blue("\n\nSticker Recieved :\n %v \n from %v\n\n", v.Message.StickerMessage.GetMimetype(), v.Info.Sender)
}
