package main

type director struct {
	builder
}

func newDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) construct() {
	d.makeTitle("Greeting")
	d.makeString("上午")
	d.makeItems([]string{
		"早上好",
		"上午好",
	})
	d.makeString("中午")
	d.makeItems([]string{
		"中午好",
	})
	d.makeString("晚上")
	d.makeItems([]string{
		"晚上好",
		"晚安",
	})
	d.output()
}
