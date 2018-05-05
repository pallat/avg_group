package average

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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

func TestSpawnGroupOfAverage(t *testing.T) {
	names := []string{"A", "WP"}

	ward := New(names)

	var i float64

	go func() {
		rand.Seed(int64(time.Now().Nanosecond()))
		for i = 0; i < 100; i++ {
			ward.Put(data{SecuritySymbol: "A", MarketCap: rand.Float64() * i})
			<-time.After(time.Duration(rand.Int63n(500)) * time.Millisecond)
		}
	}()
	go func() {
		rand.Seed(int64(time.Now().Nanosecond()))
		for i = 0; i < 100; i++ {
			ward.Put(data{SecuritySymbol: "WP", MarketCap: rand.Float64() * i})
			<-time.After(time.Duration(rand.Int63n(500)) * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		<-time.After(2 * time.Second)
		fmt.Println(ward.Get("A"))
		fmt.Println(ward.Get("WP"))
	}
}
