package view

import "Odin/model"

type WordItem struct {
	Word  string
	Count int
	// 传输前转换为具体字符串，避免客户端分不清
	Pos      string `json:"pos"`
	TalkWith string `json:"talk_with"`
	Sentence string `json:"sentence"`
	SendTime string `json:"send_time"`
}

type WordItems struct {
	Words []*WordItem
}

func FormatWordItems(data *model.WordItems) *WordItems {
	formatWordItems := WordItems{Words: []*WordItem{}}
	for _, wordItem := range *data {
		var pos string
		switch wordItem.Pos {
		case model.NOUN:
			pos = "noun"
		case model.VERB:
			pos = "verb"
		case model.ADJECTIVE:
			pos = "adjective"
		case model.PHRASE:
			pos = "phrase"
		default:
			pos = "phrase" // 默认按照短语（最不常用）的词性分类
		}
		formatWordItems.Words = append(formatWordItems.Words, &WordItem{
			Word:     wordItem.Word,
			Count:    wordItem.Count,
			Pos:      pos,
			TalkWith: wordItem.TalkWith,
			Sentence: wordItem.Sentence,
			SendTime: wordItem.SendTime,
		})
	}
	return &formatWordItems
}
