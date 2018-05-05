package average

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

type Coop struct {
	m    map[string]chan data
	ward map[string]getter
}

func (c *Coop) Put(d data) {
	c.m[d.SecuritySymbol] <- d
}

func (c *Coop) Get(n string) data {
	return c.ward[n]()
}

func New(a []string) *Coop {
	w := make(map[string]getter)
	m := make(map[string]chan data)

	for _, v := range a {
		m[v] = make(chan data)
		w[v] = initial(v, m[v])
	}

	return &Coop{
		m:    m,
		ward: w,
	}
}
