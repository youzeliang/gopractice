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

// NewBTree 创建一个BTree索引
func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32), // 这里的32是叶子节点的数量
		luck: new(sync.RWMutex),
	}
}

// 实现的话 主要是
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	// item 是一个抽象的接口，需要实现Less方法

	it := &Item{key: key, pos: pos}
	// 存储之前要对其加锁

	bt.luck.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.luck.Unlock()

	return true
}
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	// 如果是空的直接返回
	if btreeItem == nil {
		return nil
	}
	// 否则就返回、这里需要对齐进行一个转换
	return btreeItem.(*Item).pos
}

func (bt *BTree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.luck.Lock()
	oldItem := bt.tree.Delete(it)
	bt.luck.Unlock()

	// 如果删除的时候，发现没有key，就返回false 说明是无效的值
	if oldItem == nil {
		return false
	}
	return true
}
