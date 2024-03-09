package index

import (
	"github.com/google/btree"
	"goalgorithm/kv/data"
	"sync"
)

// BTree 索引，主要封装了Google的btree库
// 并发写入的时候，需要加锁，因为是线程不安全的，但并发读是安全的
type BTree struct {
	tree *btree.BTree
	luck *sync.RWMutex
}

func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	return false
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	return nil
}

func (bt *BTree) Delete(key []byte) bool {
	return false
}
