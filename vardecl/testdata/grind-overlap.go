package p

func f() {
	var x int
	{
		if b1 {
			x = 2
		} else if b2 {
			x = 4
			x = 2 * x
		} else {
			x = 3
		}
		use(x)
	}
	{
		x = 2
		use(x)
	}
}

var b1, b2 bool

func use(interface{})
