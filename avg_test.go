package main

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	var input = []data{
		{SecuritySymbol: "A", MarketCap: 100.00000},
		{SecuritySymbol: "A", MarketCap: 200.00000},
	}

	expeted := data{SecuritySymbol: "A", MarketCap: 150.00000}

	r := average(input)
	if !reflect.DeepEqual(r, expeted) {
		t.Error("wrong")
	}
}

func TestGrouping(t *testing.T) {
	var input = []data{
		{SecuritySymbol: "A", MarketCap: 100.00000},
		{SecuritySymbol: "WP", MarketCap: 200.00000},
		{SecuritySymbol: "A", MarketCap: 200.00000},
		{SecuritySymbol: "WP", MarketCap: 300.00000},
	}

	g := group(input)

	expected := map[string][]data{
		"A": []data{
			{SecuritySymbol: "A", MarketCap: 100.00000},
			{SecuritySymbol: "A", MarketCap: 200.00000},
		},
		"WP": []data{
			{SecuritySymbol: "WP", MarketCap: 200.00000},
			{SecuritySymbol: "WP", MarketCap: 300.00000},
		},
	}

	if !reflect.DeepEqual(g, expected) {
		t.Error("wrong")
	}

}

func TestAvg(t *testing.T) {
	var input = []data{
		{SecuritySymbol: "A", MarketCap: 100.00000},
		{SecuritySymbol: "WP", MarketCap: 200.00000},
		{SecuritySymbol: "A", MarketCap: 200.00000},
		{SecuritySymbol: "WP", MarketCap: 300.00000},
	}

	r := Avg(input)
	var expected = []data{
		{SecuritySymbol: "A", MarketCap: 150.00000},
		{SecuritySymbol: "WP", MarketCap: 250.00000},
	}

	if !reflect.DeepEqual(r, expected) {
		t.Error("wrong")
	}
}
