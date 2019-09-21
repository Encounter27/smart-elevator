package main

import "fmt"

type STATUS int
type DIRECTION int

const (
	UPWARD   = DIRECTION(0)
	DOWNWARD = DIRECTION(1)
	PAUSED   = DIRECTION(2)
	CLOSE    = STATUS(0)
	OPEN     = STATUS(1)
)

var status []string = []string{"CLOSE", "OPEN"}

type iLift interface {
	Direction() DIRECTION
	Pos() int
	Open()
	Close()
	Move()
	State() STATUS
	PrintStatus()
	AddTarget(n int)
}

type Lift struct {
	ID                 int
	dir                DIRECTION
	pos                int
	status             STATUS
	lowerBound         int
	upperBound         int
	ActiveTargetFloors []bool
	LowestTarget       int
	HighestTarget      int
	TotalTime          int
}

func NewLift(lb int, ub int, id int) *Lift {
	return &Lift{
		ID:                 id,
		dir:                PAUSED,
		pos:                0,
		status:             CLOSE,
		lowerBound:         lb,
		upperBound:         ub,
		ActiveTargetFloors: make([]bool, ub-lb+2),
		LowestTarget:       ub,
		HighestTarget:      lb,
		TotalTime:          0,
	}
}

func (l *Lift) Direction() DIRECTION {
	return l.dir
}

func (l *Lift) Pos() int {
	return l.pos
}
func (l *Lift) Open() {
	l.TotalTime++
	l.status = OPEN
	fmt.Printf("LIFT %d OPENS", l.ID)
}
func (l *Lift) Close() {
	l.status = CLOSE
}

func (l *Lift) PrintStatus() {
	fmt.Printf("| LIFT %d --> %d(%s) | ", l.ID, l.pos, status[l.status])
}

func (l *Lift) Move() {
	l.TotalTime++
	switch l.dir {
	case UPWARD:
		if l.pos < l.HighestTarget {
			l.pos++
			if l.ActiveTargetFloors[l.pos] == true {
				l.ActiveTargetFloors[l.pos] = false
				l.Open()
			} else {
				l.ActiveTargetFloors[l.pos] = false
				l.Close()
				l.dir = PAUSED
				l.HighestTarget = -1
			}
		}
	case DOWNWARD:
		if l.pos > l.LowestTarget {
			l.pos--
			if l.ActiveTargetFloors[l.pos] == true {
				l.ActiveTargetFloors[l.pos] = false
				l.Open()
			} else {
				l.ActiveTargetFloors[l.pos] = false
				l.Close()
				l.dir = PAUSED
			}
		}
	case PAUSED:
		if l.LowestTarget < l.pos && l.pos < l.HighestTarget {
			if l.pos-l.LowestTarget < l.HighestTarget-l.pos {
				l.dir = DOWNWARD
				l.pos--
			} else {
				l.dir = UPWARD
				l.pos++
			}
		} else if l.pos <= l.LowestTarget {

		} else if l.pos >= l.HighestTarget {

		}
	}
}

func (l *Lift) AddTarget(n int) {
	if l.LowestTarget > n {
		l.LowestTarget = n
	}

	if l.HighestTarget < n {
		l.HighestTarget = n
	}
	l.ActiveTargetFloors[n] = true
}

func (l *Lift) Calculate(from int, to int) int {
	count := 0

	// logic
	var d DIRECTION
	if from < to {
		d = UPWARD
	} else {
		d = DOWNWARD
	}

	switch {
	case d == UPWARD, l.dir == UPWARD:
		return min(to, l.HighestTarget) - max(from, l.LowestTarget)

	case d == UPWARD, l.dir == DOWNWARD:
		return l.pos - l.LowestTarget + l.HighestTarget - from

	case d == DOWNWARD, l.dir == DOWNWARD:
		return min(to, l.HighestTarget) - max(from, l.LowestTarget)

	case d == DOWNWARD, l.dir == UPWARD:
		return l.HighestTarget - l.pos + l.HighestTarget - to
	}

	return count
}
