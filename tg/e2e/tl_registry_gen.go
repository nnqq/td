// Code generated by gotdgen, DO NOT EDIT.

package e2e

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/nnqq/td/bin"
	"github.com/nnqq/td/tdp"
	"github.com/nnqq/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// TypesMap returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		0xa8509bda: "int#a8509bda",
		0x22076cba: "long#22076cba",
		0x2210c154: "double#2210c154",
		0xb5286e24: "string#b5286e24",
		0xe937bb82: "bytes#e937bb82",
		0xbc799737: "boolFalse#bc799737",
		0x997275b5: "boolTrue#997275b5",
		0x3fedd339: "true#3fedd339",
		0x1f814f1f: "decryptedMessage8#1f814f1f",
		0xaa48327d: "decryptedMessageService8#aa48327d",
		0x89f5c4a:  "decryptedMessageMediaEmpty#89f5c4a",
		0x32798a8c: "decryptedMessageMediaPhoto23#32798a8c",
		0x4cee6ef3: "decryptedMessageMediaVideo8#4cee6ef3",
		0x35480a59: "decryptedMessageMediaGeoPoint#35480a59",
		0x588a0a97: "decryptedMessageMediaContact#588a0a97",
		0xa1733aec: "decryptedMessageActionSetMessageTTL#a1733aec",
		0xb095434b: "decryptedMessageMediaDocument23#b095434b",
		0x6080758f: "decryptedMessageMediaAudio8#6080758f",
		0xc4f40be:  "decryptedMessageActionReadMessages#c4f40be",
		0x65614304: "decryptedMessageActionDeleteMessages#65614304",
		0x8ac1f475: "decryptedMessageActionScreenshotMessages#8ac1f475",
		0x6719e45c: "decryptedMessageActionFlushHistory#6719e45c",
		0x204d3878: "decryptedMessage23#204d3878",
		0x73164160: "decryptedMessageService#73164160",
		0x524a415d: "decryptedMessageMediaVideo23#524a415d",
		0x57e0a9cb: "decryptedMessageMediaAudio#57e0a9cb",
		0x1be31789: "decryptedMessageLayer#1be31789",
		0x16bf744e: "sendMessageTypingAction#16bf744e",
		0xfd5ec8f5: "sendMessageCancelAction#fd5ec8f5",
		0xa187d66f: "sendMessageRecordVideoAction#a187d66f",
		0x92042ff7: "sendMessageUploadVideoAction#92042ff7",
		0xd52f73f7: "sendMessageRecordAudioAction#d52f73f7",
		0xe6ac8a6f: "sendMessageUploadAudioAction#e6ac8a6f",
		0x990a3c1a: "sendMessageUploadPhotoAction#990a3c1a",
		0x8faee98e: "sendMessageUploadDocumentAction#8faee98e",
		0x176f8ba1: "sendMessageGeoLocationAction#176f8ba1",
		0x628cbc6f: "sendMessageChooseContactAction#628cbc6f",
		0x511110b0: "decryptedMessageActionResend#511110b0",
		0xf3048883: "decryptedMessageActionNotifyLayer#f3048883",
		0xccb27641: "decryptedMessageActionTyping#ccb27641",
		0xf3c9611b: "decryptedMessageActionRequestKey#f3c9611b",
		0x6fe1735b: "decryptedMessageActionAcceptKey#6fe1735b",
		0xdd05ec6b: "decryptedMessageActionAbortKey#dd05ec6b",
		0xec2e0b9b: "decryptedMessageActionCommitKey#ec2e0b9b",
		0xa82fdd63: "decryptedMessageActionNoop#a82fdd63",
		0x6c37c15c: "documentAttributeImageSize#6c37c15c",
		0x11b58939: "documentAttributeAnimated#11b58939",
		0xfb0a5727: "documentAttributeSticker23#fb0a5727",
		0x5910cccb: "documentAttributeVideo#5910cccb",
		0x51448e5:  "documentAttributeAudio23#51448e5",
		0x15590068: "documentAttributeFilename#15590068",
		0xe17e23c:  "photoSizeEmpty#e17e23c",
		0x77bfb61b: "photoSize#77bfb61b",
		0xe9a734fa: "photoCachedSize#e9a734fa",
		0x7c596b46: "fileLocationUnavailable#7c596b46",
		0x53d69076: "fileLocation#53d69076",
		0xfa95b0dd: "decryptedMessageMediaExternalDocument#fa95b0dd",
		0xded218e0: "documentAttributeAudio45#ded218e0",
		0x36b091de: "decryptedMessage46#36b091de",
		0xf1fa8d78: "decryptedMessageMediaPhoto#f1fa8d78",
		0x970c8c0e: "decryptedMessageMediaVideo#970c8c0e",
		0x7afe8ae2: "decryptedMessageMediaDocument#7afe8ae2",
		0x3a556302: "documentAttributeSticker#3a556302",
		0x9852f9c6: "documentAttributeAudio#9852f9c6",
		0xbb92ba95: "messageEntityUnknown#bb92ba95",
		0xfa04579d: "messageEntityMention#fa04579d",
		0x6f635b0d: "messageEntityHashtag#6f635b0d",
		0x6cef8ac7: "messageEntityBotCommand#6cef8ac7",
		0x6ed02538: "messageEntityUrl#6ed02538",
		0x64e475c2: "messageEntityEmail#64e475c2",
		0xbd610bc9: "messageEntityBold#bd610bc9",
		0x826f8b60: "messageEntityItalic#826f8b60",
		0x28a20571: "messageEntityCode#28a20571",
		0x73924be0: "messageEntityPre#73924be0",
		0x76a6d327: "messageEntityTextUrl#76a6d327",
		0x352dca58: "messageEntityMentionName#352dca58",
		0x9b69e34b: "messageEntityPhone#9b69e34b",
		0x4c4e743f: "messageEntityCashtag#4c4e743f",
		0x761e6af4: "messageEntityBankCard#761e6af4",
		0x861cc8a0: "inputStickerSetShortName#861cc8a0",
		0xffb62b95: "inputStickerSetEmpty#ffb62b95",
		0x8a0df56f: "decryptedMessageMediaVenue#8a0df56f",
		0xe50511d8: "decryptedMessageMediaWebPage#e50511d8",
		0x88f27fbc: "sendMessageRecordRoundAction#88f27fbc",
		0xbb718624: "sendMessageUploadRoundAction#bb718624",
		0xef02ce6:  "documentAttributeVideo66#ef02ce6",
		0x91cc4674: "decryptedMessage#91cc4674",
		0x9c4e7e8b: "messageEntityUnderline#9c4e7e8b",
		0xbf0693d4: "messageEntityStrike#bf0693d4",
		0x20df5d0:  "messageEntityBlockquote#20df5d0",
		0xc8357709: "test.dummyFunction#c8357709",
	}
}

// NamesMap returns mapping from type names to TL type ids.
func NamesMap() map[string]uint32 {
	return map[string]uint32{
		"int":                                      0xa8509bda,
		"long":                                     0x22076cba,
		"double":                                   0x2210c154,
		"string":                                   0xb5286e24,
		"bytes":                                    0xe937bb82,
		"boolFalse":                                0xbc799737,
		"boolTrue":                                 0x997275b5,
		"true":                                     0x3fedd339,
		"decryptedMessage8":                        0x1f814f1f,
		"decryptedMessageService8":                 0xaa48327d,
		"decryptedMessageMediaEmpty":               0x89f5c4a,
		"decryptedMessageMediaPhoto23":             0x32798a8c,
		"decryptedMessageMediaVideo8":              0x4cee6ef3,
		"decryptedMessageMediaGeoPoint":            0x35480a59,
		"decryptedMessageMediaContact":             0x588a0a97,
		"decryptedMessageActionSetMessageTTL":      0xa1733aec,
		"decryptedMessageMediaDocument23":          0xb095434b,
		"decryptedMessageMediaAudio8":              0x6080758f,
		"decryptedMessageActionReadMessages":       0xc4f40be,
		"decryptedMessageActionDeleteMessages":     0x65614304,
		"decryptedMessageActionScreenshotMessages": 0x8ac1f475,
		"decryptedMessageActionFlushHistory":       0x6719e45c,
		"decryptedMessage23":                       0x204d3878,
		"decryptedMessageService":                  0x73164160,
		"decryptedMessageMediaVideo23":             0x524a415d,
		"decryptedMessageMediaAudio":               0x57e0a9cb,
		"decryptedMessageLayer":                    0x1be31789,
		"sendMessageTypingAction":                  0x16bf744e,
		"sendMessageCancelAction":                  0xfd5ec8f5,
		"sendMessageRecordVideoAction":             0xa187d66f,
		"sendMessageUploadVideoAction":             0x92042ff7,
		"sendMessageRecordAudioAction":             0xd52f73f7,
		"sendMessageUploadAudioAction":             0xe6ac8a6f,
		"sendMessageUploadPhotoAction":             0x990a3c1a,
		"sendMessageUploadDocumentAction":          0x8faee98e,
		"sendMessageGeoLocationAction":             0x176f8ba1,
		"sendMessageChooseContactAction":           0x628cbc6f,
		"decryptedMessageActionResend":             0x511110b0,
		"decryptedMessageActionNotifyLayer":        0xf3048883,
		"decryptedMessageActionTyping":             0xccb27641,
		"decryptedMessageActionRequestKey":         0xf3c9611b,
		"decryptedMessageActionAcceptKey":          0x6fe1735b,
		"decryptedMessageActionAbortKey":           0xdd05ec6b,
		"decryptedMessageActionCommitKey":          0xec2e0b9b,
		"decryptedMessageActionNoop":               0xa82fdd63,
		"documentAttributeImageSize":               0x6c37c15c,
		"documentAttributeAnimated":                0x11b58939,
		"documentAttributeSticker23":               0xfb0a5727,
		"documentAttributeVideo":                   0x5910cccb,
		"documentAttributeAudio23":                 0x51448e5,
		"documentAttributeFilename":                0x15590068,
		"photoSizeEmpty":                           0xe17e23c,
		"photoSize":                                0x77bfb61b,
		"photoCachedSize":                          0xe9a734fa,
		"fileLocationUnavailable":                  0x7c596b46,
		"fileLocation":                             0x53d69076,
		"decryptedMessageMediaExternalDocument":    0xfa95b0dd,
		"documentAttributeAudio45":                 0xded218e0,
		"decryptedMessage46":                       0x36b091de,
		"decryptedMessageMediaPhoto":               0xf1fa8d78,
		"decryptedMessageMediaVideo":               0x970c8c0e,
		"decryptedMessageMediaDocument":            0x7afe8ae2,
		"documentAttributeSticker":                 0x3a556302,
		"documentAttributeAudio":                   0x9852f9c6,
		"messageEntityUnknown":                     0xbb92ba95,
		"messageEntityMention":                     0xfa04579d,
		"messageEntityHashtag":                     0x6f635b0d,
		"messageEntityBotCommand":                  0x6cef8ac7,
		"messageEntityUrl":                         0x6ed02538,
		"messageEntityEmail":                       0x64e475c2,
		"messageEntityBold":                        0xbd610bc9,
		"messageEntityItalic":                      0x826f8b60,
		"messageEntityCode":                        0x28a20571,
		"messageEntityPre":                         0x73924be0,
		"messageEntityTextUrl":                     0x76a6d327,
		"messageEntityMentionName":                 0x352dca58,
		"messageEntityPhone":                       0x9b69e34b,
		"messageEntityCashtag":                     0x4c4e743f,
		"messageEntityBankCard":                    0x761e6af4,
		"inputStickerSetShortName":                 0x861cc8a0,
		"inputStickerSetEmpty":                     0xffb62b95,
		"decryptedMessageMediaVenue":               0x8a0df56f,
		"decryptedMessageMediaWebPage":             0xe50511d8,
		"sendMessageRecordRoundAction":             0x88f27fbc,
		"sendMessageUploadRoundAction":             0xbb718624,
		"documentAttributeVideo66":                 0xef02ce6,
		"decryptedMessage":                         0x91cc4674,
		"messageEntityUnderline":                   0x9c4e7e8b,
		"messageEntityStrike":                      0xbf0693d4,
		"messageEntityBlockquote":                  0x20df5d0,
		"test.dummyFunction":                       0xc8357709,
	}
}

// TypesConstructorMap maps type ids to constructors.
func TypesConstructorMap() map[uint32]func() bin.Object {
	return map[uint32]func() bin.Object{
		0xa8509bda: func() bin.Object { return &Int{} },
		0x22076cba: func() bin.Object { return &Long{} },
		0x2210c154: func() bin.Object { return &Double{} },
		0xb5286e24: func() bin.Object { return &String{} },
		0xe937bb82: func() bin.Object { return &Bytes{} },
		0xbc799737: func() bin.Object { return &BoolFalse{} },
		0x997275b5: func() bin.Object { return &BoolTrue{} },
		0x3fedd339: func() bin.Object { return &True{} },
		0x1f814f1f: func() bin.Object { return &DecryptedMessage8{} },
		0xaa48327d: func() bin.Object { return &DecryptedMessageService8{} },
		0x89f5c4a:  func() bin.Object { return &DecryptedMessageMediaEmpty{} },
		0x32798a8c: func() bin.Object { return &DecryptedMessageMediaPhoto23{} },
		0x4cee6ef3: func() bin.Object { return &DecryptedMessageMediaVideo8{} },
		0x35480a59: func() bin.Object { return &DecryptedMessageMediaGeoPoint{} },
		0x588a0a97: func() bin.Object { return &DecryptedMessageMediaContact{} },
		0xa1733aec: func() bin.Object { return &DecryptedMessageActionSetMessageTTL{} },
		0xb095434b: func() bin.Object { return &DecryptedMessageMediaDocument23{} },
		0x6080758f: func() bin.Object { return &DecryptedMessageMediaAudio8{} },
		0xc4f40be:  func() bin.Object { return &DecryptedMessageActionReadMessages{} },
		0x65614304: func() bin.Object { return &DecryptedMessageActionDeleteMessages{} },
		0x8ac1f475: func() bin.Object { return &DecryptedMessageActionScreenshotMessages{} },
		0x6719e45c: func() bin.Object { return &DecryptedMessageActionFlushHistory{} },
		0x204d3878: func() bin.Object { return &DecryptedMessage23{} },
		0x73164160: func() bin.Object { return &DecryptedMessageService{} },
		0x524a415d: func() bin.Object { return &DecryptedMessageMediaVideo23{} },
		0x57e0a9cb: func() bin.Object { return &DecryptedMessageMediaAudio{} },
		0x1be31789: func() bin.Object { return &DecryptedMessageLayer{} },
		0x16bf744e: func() bin.Object { return &SendMessageTypingAction{} },
		0xfd5ec8f5: func() bin.Object { return &SendMessageCancelAction{} },
		0xa187d66f: func() bin.Object { return &SendMessageRecordVideoAction{} },
		0x92042ff7: func() bin.Object { return &SendMessageUploadVideoAction{} },
		0xd52f73f7: func() bin.Object { return &SendMessageRecordAudioAction{} },
		0xe6ac8a6f: func() bin.Object { return &SendMessageUploadAudioAction{} },
		0x990a3c1a: func() bin.Object { return &SendMessageUploadPhotoAction{} },
		0x8faee98e: func() bin.Object { return &SendMessageUploadDocumentAction{} },
		0x176f8ba1: func() bin.Object { return &SendMessageGeoLocationAction{} },
		0x628cbc6f: func() bin.Object { return &SendMessageChooseContactAction{} },
		0x511110b0: func() bin.Object { return &DecryptedMessageActionResend{} },
		0xf3048883: func() bin.Object { return &DecryptedMessageActionNotifyLayer{} },
		0xccb27641: func() bin.Object { return &DecryptedMessageActionTyping{} },
		0xf3c9611b: func() bin.Object { return &DecryptedMessageActionRequestKey{} },
		0x6fe1735b: func() bin.Object { return &DecryptedMessageActionAcceptKey{} },
		0xdd05ec6b: func() bin.Object { return &DecryptedMessageActionAbortKey{} },
		0xec2e0b9b: func() bin.Object { return &DecryptedMessageActionCommitKey{} },
		0xa82fdd63: func() bin.Object { return &DecryptedMessageActionNoop{} },
		0x6c37c15c: func() bin.Object { return &DocumentAttributeImageSize{} },
		0x11b58939: func() bin.Object { return &DocumentAttributeAnimated{} },
		0xfb0a5727: func() bin.Object { return &DocumentAttributeSticker23{} },
		0x5910cccb: func() bin.Object { return &DocumentAttributeVideo{} },
		0x51448e5:  func() bin.Object { return &DocumentAttributeAudio23{} },
		0x15590068: func() bin.Object { return &DocumentAttributeFilename{} },
		0xe17e23c:  func() bin.Object { return &PhotoSizeEmpty{} },
		0x77bfb61b: func() bin.Object { return &PhotoSize{} },
		0xe9a734fa: func() bin.Object { return &PhotoCachedSize{} },
		0x7c596b46: func() bin.Object { return &FileLocationUnavailable{} },
		0x53d69076: func() bin.Object { return &FileLocation{} },
		0xfa95b0dd: func() bin.Object { return &DecryptedMessageMediaExternalDocument{} },
		0xded218e0: func() bin.Object { return &DocumentAttributeAudio45{} },
		0x36b091de: func() bin.Object { return &DecryptedMessage46{} },
		0xf1fa8d78: func() bin.Object { return &DecryptedMessageMediaPhoto{} },
		0x970c8c0e: func() bin.Object { return &DecryptedMessageMediaVideo{} },
		0x7afe8ae2: func() bin.Object { return &DecryptedMessageMediaDocument{} },
		0x3a556302: func() bin.Object { return &DocumentAttributeSticker{} },
		0x9852f9c6: func() bin.Object { return &DocumentAttributeAudio{} },
		0xbb92ba95: func() bin.Object { return &MessageEntityUnknown{} },
		0xfa04579d: func() bin.Object { return &MessageEntityMention{} },
		0x6f635b0d: func() bin.Object { return &MessageEntityHashtag{} },
		0x6cef8ac7: func() bin.Object { return &MessageEntityBotCommand{} },
		0x6ed02538: func() bin.Object { return &MessageEntityURL{} },
		0x64e475c2: func() bin.Object { return &MessageEntityEmail{} },
		0xbd610bc9: func() bin.Object { return &MessageEntityBold{} },
		0x826f8b60: func() bin.Object { return &MessageEntityItalic{} },
		0x28a20571: func() bin.Object { return &MessageEntityCode{} },
		0x73924be0: func() bin.Object { return &MessageEntityPre{} },
		0x76a6d327: func() bin.Object { return &MessageEntityTextURL{} },
		0x352dca58: func() bin.Object { return &MessageEntityMentionName{} },
		0x9b69e34b: func() bin.Object { return &MessageEntityPhone{} },
		0x4c4e743f: func() bin.Object { return &MessageEntityCashtag{} },
		0x761e6af4: func() bin.Object { return &MessageEntityBankCard{} },
		0x861cc8a0: func() bin.Object { return &InputStickerSetShortName{} },
		0xffb62b95: func() bin.Object { return &InputStickerSetEmpty{} },
		0x8a0df56f: func() bin.Object { return &DecryptedMessageMediaVenue{} },
		0xe50511d8: func() bin.Object { return &DecryptedMessageMediaWebPage{} },
		0x88f27fbc: func() bin.Object { return &SendMessageRecordRoundAction{} },
		0xbb718624: func() bin.Object { return &SendMessageUploadRoundAction{} },
		0xef02ce6:  func() bin.Object { return &DocumentAttributeVideo66{} },
		0x91cc4674: func() bin.Object { return &DecryptedMessage{} },
		0x9c4e7e8b: func() bin.Object { return &MessageEntityUnderline{} },
		0xbf0693d4: func() bin.Object { return &MessageEntityStrike{} },
		0x20df5d0:  func() bin.Object { return &MessageEntityBlockquote{} },
		0xc8357709: func() bin.Object { return &TestDummyFunctionRequest{} },
	}
}
