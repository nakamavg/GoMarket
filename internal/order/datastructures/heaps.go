package datastructures

import (
	"market/internal/order/model"
)

type BuyHeap []*model.Order

func (h BuyHeap) Len() int           { return len(h) }
func (h BuyHeap) Less(i, j int) bool { return h[i].LimitPrice > h[j].LimitPrice }
func (h BuyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BuyHeap) Push(x interface{}) { *h = append(*h, x.(*model.Order)) }
func (h *BuyHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type SellHeap []*model.Order

func (h SellHeap) Len() int           { return len(h) }
func (h SellHeap) Less(i, j int) bool { return h[i].LimitPrice < h[j].LimitPrice }
func (h SellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SellHeap) Push(x interface{}) { *h = append(*h, x.(*model.Order)) }
func (h *SellHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}