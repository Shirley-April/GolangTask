package models

type Pair struct {
	Key interface{}
	Value int
}

type PairList []Pair

func (p PairList) Len() int  {
	return len(p)
}

func (p PairList) Less(index1, index2 int)bool  {
	return p[index1].Value < p[index2].Value
}

func (p PairList) Swap(index1, index2 int)  {
	p[index1], p[index2] = p[index1], p[index2]
}