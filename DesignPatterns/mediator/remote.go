package main

// implements mediator
type remote struct {
	c *computer
	o *oven
	f *fan
	r *refrigerator
}

func newRemote() *remote {
	return &remote{}
}

func (r *remote) config() (*computer, *oven) {
	c := newComputer()
	c.setMediator(r)
	r.c = c
	o := newOven()
	o.setMediator(r)
	r.o = o
	f := newFan()
	f.setMediator(r)
	r.f = f
	rf := newRefrigerator()
	rf.setMediator(r)
	r.r = rf
	return c, o
}

func (r *remote) turnedOven(enabled bool) {
	if enabled {
		r.c.setStatus(false)
		r.r.setStatus(false)
	} else {
		r.r.setStatus(true)
	}
}

func (r *remote) turnedComputer(enabled bool) {
	r.f.setStatus(enabled)
}
