package main

import (
	"fmt"
	"strconv"
)

type Controller struct {
	lifts       []*Lift
	time        int
	totalFloors int
	totalLifts  int
}

func NewController(lifts int, floor int) *Controller {
	c := new(Controller)
	c.lifts = make([]*Lift, lifts)
	c.totalFloors = floor
	c.time = -1

	for i := 0; i < lifts; i++ {
		c.lifts[i] = NewLift(0, floor, i+1)
	}

	return c
}

func (c *Controller) AssignLift() {

}

func (c *Controller) ProcessRequest(args []string) {
	if len(args) != 0 {

		requests := make([]int, len(args))

		for i := 0; i < len(args); i++ {
			requests[i], _ = strconv.Atoi(args[i])
		}

		for i := 0; i < len(args); i = i + 2 {
			from := requests[i]
			to := requests[i+1]

			estimate := 10000 // some max value
			choice := 0

			for e, l := range c.lifts {
				count := l.Calculate(from, to)

				if count < estimate {
					estimate = count
					choice = e
				}
			}

			c.lifts[choice].AddTarget(from)
			c.lifts[choice].AddTarget(to)
		}
	}
	for _, l := range c.lifts {
		l.Move()
	}

	//if args == nil {
	fmt.Printf("T = %d --------\n", c.time)
	c.StatusAll()
	//}

	c.time++
}

func (c *Controller) StatusAll() {
	for i := 0; i < len(c.lifts); i++ {
		c.lifts[i].PrintStatus()
	}
	fmt.Printf("\n")
}
