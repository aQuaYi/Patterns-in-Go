package main

type manager struct {
	showcase map[string]product
}

func newManager() *manager {
	return &manager{
		showcase: make(map[string]product, 16),
	}
}

func (m *manager) register(name string, p product) {
	m.showcase[name] = p
}

func (m *manager) create(name string) product {
	p := m.showcase[name]
	return p.createClone()
}
