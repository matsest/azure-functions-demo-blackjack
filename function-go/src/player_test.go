package main

import (
	"reflect"
	"testing"
)

func TestPlayer_drawCards(t *testing.T) {
	type fields struct {
		Name  string
		Hand  []Card
		Value int
	}
	type args struct {
		cards  *[]Card
		number int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &Player{
				Name:  tt.fields.Name,
				Hand:  tt.fields.Hand,
				Value: tt.fields.Value,
			}
			player.drawCards(tt.args.cards, tt.args.number)
		})
	}
}

func Test_newPlayer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPlayer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
