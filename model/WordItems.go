package model

type DataAmountType int

const (
	INTENSIVE_DATA DataAmountType = 0
	MODERATE_DATA  DataAmountType = 1
	SPARSE_DATA    DataAmountType = 2
)

type WordItems []*WordItem

func (wordItems WordItem) Len() int {
	return wordItems.Len()
}

func (wordItems WordItems) Less(i, j int) bool {
	return wordItems[i].Count > wordItems[j].Count
}

func (wordItems WordItems) Swap(i, j int) {
	wordItems[i], wordItems[j] = wordItems[j], wordItems[i]
}
