package main

func newChStream() map[string]chan data {
	return map[string]chan data{
		"A":  make(chan data),
		"WP": make(chan data),
	}
}

type getter func() data
type grab func() getter

func initial(s string, ch chan data) getter {
	name := s
	var sum float64 = 0
	var total float64 = 0

	go func(dch chan data) {
		for {
			v := <-dch
			sum += v.MarketCap
			total++
		}
	}(ch)

	return func() data {
		return data{SecuritySymbol: name, MarketCap: sum / total}
	}
}
