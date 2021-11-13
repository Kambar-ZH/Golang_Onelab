package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func() {
		for num := range(c.Input) {
			c.Output <- num * num
		}
		defer close(c.Output)
	} ()
}
