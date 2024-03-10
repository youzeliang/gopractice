package index

import (
	"bytes"
	"github.com/google/btree"
	"goalgorithm/kv/data"
)

// Indexer 抽象索引接口， 后续如果想接入其他的数据结构，只需要实现这个接口即可
type Indexer interface {
	// Put 向索引中存储 key 对应的数据位置信息
	Put(key []byte, pos *data.LogRecordPos) bool
	// Get 从索引中获取 key 对应的数据位置信息
	Get(key []byte) *data.LogRecordPos
	// Delete 从索引中删除 key 对应的数据位置信息
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (i *Item) Less(than btree.Item) bool {
	return bytes.Compare(i.key, than.(*Item).key) == -1
}
