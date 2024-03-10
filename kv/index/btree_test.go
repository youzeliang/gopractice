package index

import "testing"

func TestBTree_Put(t *testing.T) {
	type fields struct {
		tree *btree.BTree
		luck *sync.RWMutex
	}
	type args struct {
		key []byte
		pos *data.LogRecordPos
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BTree{
				tree: tt.fields.tree,
				luck: tt.fields.luck,
			}
			if got := bt.Put(tt.args.key, tt.args.pos); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
