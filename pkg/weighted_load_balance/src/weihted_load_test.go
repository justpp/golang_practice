package src

import (
	"testing"
)

// 没用的测试

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		{"一个不成熟的benchmark"},
	}
	a := &WeightRoundBalance{}
	a.Add("127.0.0.1:99", 2)
	a.Add("127.0.0.1:999", 4)
	a.Add("127.0.0.1:9999", 5)
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a.Next()
			}
		})
	}
}

func TestWeightRoundBalance_Add(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []*WeightNode
	}
	type args struct {
		addr   string
		weight int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "1", args: args{addr: "127.0.0.1:9090", weight: 2}},
		{name: "2", args: args{addr: "127.0.0.1:9090", weight: 3}},
		{name: "3", args: args{addr: "127.0.0.1:9090", weight: 2}},
		{name: "4", args: args{addr: "127.0.0.1:9090", weight: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeightRoundBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			w.Add(tt.args.addr, tt.args.weight)
		})
	}
}

func TestWeightRoundBalance_Get(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []*WeightNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "1", fields: fields{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeightRoundBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			if got := w.Get(); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeightRoundBalance_Next(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []*WeightNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "1", fields: fields{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeightRoundBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			if got := w.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
