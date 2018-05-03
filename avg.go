package main

func main() {

}

type data struct {
	SecuritySymbol string  `json:"securitySymbol"`
	MarketCap      float64 `json:"marketCap"`
}

func average(in []data) data {
	var sum float64 = 0
	for _, v := range in {
		sum += v.MarketCap
	}
	return data{SecuritySymbol: in[0].SecuritySymbol, MarketCap: sum / float64(len(in))}
}

func group(in []data) map[string][]data {
	m := map[string][]data{}

	for _, v := range in {
		if _, ok := m[v.SecuritySymbol]; !ok {
			m[v.SecuritySymbol] = []data{}
		}

		m[v.SecuritySymbol] = append(m[v.SecuritySymbol], v)
	}

	return m
}

func Avg(in []data) []data {
	out := []data{}
	g := group(in)

	for _, v := range g {
		out = append(out, average(v))
	}
	return out
}
