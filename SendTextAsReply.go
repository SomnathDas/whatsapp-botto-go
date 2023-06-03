package main

import (
	"context"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func SendTextAsReply(WAClient *whatsmeow.Client, v *events.Message, aiResponse string) {
	// Showing [bot... is typing....] in whatsapp until message is sent
	WAClient.SendChatPresence(v.Info.Chat, types.ChatPresenceComposing, types.ChatPresenceMediaText)

	WAClient.SendMessage(context.Background(), v.Info.Chat, &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: proto.String(aiResponse),
			// Providing ContextInfo to send this message as a "reply"
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      &v.Info.ID,
				QuotedMessage: v.Message.ExtendedTextMessage.ContextInfo.GetQuotedMessage(),
				Participant:   proto.String(v.Info.MessageSource.Sender.String()),
			},
		},
	})
}
