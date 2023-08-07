package load_balance

import (
	"reflect"
	"testing"
)

func TestRoundRobinBalance_Add(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []string
	}
	type args struct {
		params []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"t1", fields{curIndex: 0, rss: nil}, args{[]string{"ss1", "ss2", "ss3", "ss1"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoundRobinBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			if err := r.Add(tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoundRobinBalance_Get(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []string
	}
	ssList := []string{"ss1", "ss2", "ss3"}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"t1", fields{0, ssList}, ssList, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoundRobinBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			var gots []string
			for i := 0; i < len(tt.want); i++ {
				got, err := r.Get()
				if err != nil {
					return
				}
				gots = append(gots, got)
			}
			if !reflect.DeepEqual(gots, tt.want) {
				t.Errorf("Get() got = %v, want %v", gots, tt.want)
			}
		})
	}
}

func TestRoundRobinBalance_Next(t *testing.T) {
	type fields struct {
		curIndex int
		rss      []string
	}
	ssList := []string{"ss1", "ss2", "ss3"}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"t1", fields{0, ssList}, ssList},
		{"t2", fields{0, ssList}, []string{"ss1", "ss2", "ss3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoundRobinBalance{
				curIndex: tt.fields.curIndex,
				rss:      tt.fields.rss,
			}
			var gets []string

			for i := 0; i < len(tt.want); i++ {
				gets = append(gets, r.Next())
			}

			if !reflect.DeepEqual(gets, tt.want) {
				t.Errorf("Next() = %v, want %v", gets, tt.want)
			}
		})
	}
}
