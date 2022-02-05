package weight

// this package implement weighted round-robin load balancing algorithm
// ---
// first, it will increase current_weight by its weight
// then it will choose backend server that have biggest current_weight
// after that, selected backend current_weight will be reduced by total_weight

// Backend hold information about backend server
type Backend struct {
	Name          string
	Weight        int
	CurrentWeight int
}

type Weight struct {
	Items []*Backend
}

// NewWeight returns Wight struct
func NewWeight() *Weight {
	return &Weight{}
}

// AddWeight add server and weight information
func (w *Weight) AddWeight(name string, weight int) {
	w.Items = append(w.Items, &Backend{
		Name:   name,
		Weight: weight,
	})
}

// Next returns next selected server
func (w *Weight) Next() string {
	// update current_weight by adding it with weight
	for i := 0; i < len(w.Items); i++ {
		w.Items[i].CurrentWeight += w.Items[i].Weight
	}

	return w.nextSelectedServer()
}

// nextSelectedServer returns selected server and update current_weight
func (w *Weight) nextSelectedServer() string {
	var server string
	var n, index int

	for i, b := range w.Items {
		if n < b.CurrentWeight {
			n = b.CurrentWeight
			server = b.Name
			index = i
		}
	}

	w.updateEffectiveWeight(index)

	return server
}

// updateEffectiveWeight update current_weight by reducing it with total_weight
// current_weight - total_weight
func (w *Weight) updateEffectiveWeight(index int) {
	for i := 0; i < len(w.Items); i++ {
		if i == index {
			w.Items[i].CurrentWeight -= w.getTotalWeight()
		}
	}
}

// getTotalWeight get total weight of all weighted server
func (w *Weight) getTotalWeight() int {
	totalWeight := 0

	for _, t := range w.Items {
		totalWeight += t.Weight
	}

	return totalWeight
}
