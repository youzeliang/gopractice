package index

import "goalgorithm/kv/data"

// Indexer 抽象索引接口， 后续如果想接入其他的数据结构，只需要实现这个接口即可
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}
