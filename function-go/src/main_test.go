package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPlayBlackJack(t *testing.T) {
	type args struct {
		name1    string
		name2    string
		cardDeck []Card
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlayBlackJack(tt.args.name1, tt.args.name2, tt.args.cardDeck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayBlackJack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_play(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			play(tt.args.w, tt.args.r)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
