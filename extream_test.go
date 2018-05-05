package average

import (
	"fmt"
	"testing"
)

func TestExtreamFunc(t *testing.T) {
	m := newChStream()

	a := initial("A", m["A"])
	wp := initial("WP", m["WP"])

	var i float64
	for i = 0; i < 100; i++ {
		m["A"] <- data{SecuritySymbol: "A", MarketCap: 100.00 * i}
		m["WP"] <- data{SecuritySymbol: "WP", MarketCap: 110.00 * i}
	}

	fmt.Println(a())
	fmt.Println(wp())
}
