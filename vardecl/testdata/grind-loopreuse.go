package p

func f() {
	var i int

	{
		for {
			use(i)
			i++
		}
	}
}

func use(interface{})
