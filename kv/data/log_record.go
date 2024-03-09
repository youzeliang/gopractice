package data

// 数据内存索引,主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件ID,表示将数据存储到了哪个文件中
	Offset int64  // 偏移，表示数据存储在文件中的什么位置
}
