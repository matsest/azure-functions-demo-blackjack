package main

import (
	"reflect"
	"testing"
)

func Test_cardValue(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want int
	}{
		{
			name: "jack",
			card: Card{"SPADES", "J"},
			want: 10,
		},
		{
			name: "ace",
			card: Card{"SPADES", "A"},
			want: 11,
		},
		{
			name: "regular",
			card: Card{"HEART", "2"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardValue(tt.card); got != tt.want {
				t.Errorf("cardValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abbreviateCards(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abbreviateCards(tt.args.cards...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("abbreviateCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardsValues(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardsValues(tt.args.cards...); got != tt.want {
				t.Errorf("cardsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadCards(t *testing.T) {
	// TODO: Must mock
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want []Card
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downloadCards(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drawCardsFromTop(t *testing.T) {
	type args struct {
		cards  []Card
		number int
	}
	tests := []struct {
		name  string
		args  args
		want  []Card
		want1 []Card
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := drawCardsFromTop(tt.args.cards, tt.args.number)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawCardsFromTop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("drawCardsFromTop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
