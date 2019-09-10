package model

// 词性
type PartOfSpeech int

const (
	NOUN      PartOfSpeech = 0 // 名词
	VERB      PartOfSpeech = 1 // 动词
	ADJECTIVE PartOfSpeech = 2 // 形容词
	PHRASE    PartOfSpeech = 3 // 短语
)

type WordItem struct {
	Word     string
	Count    int
	Pos      PartOfSpeech
	TalkWith string // 聊天对象
	Sentence string // 所在句子
	SendTime string // 发送时间
}
