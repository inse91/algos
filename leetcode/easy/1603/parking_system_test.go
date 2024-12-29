package _1603_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/design-parking-system/

func Test(t *testing.T) {
	ps := Constructor(1, 1, 0)
	assert.True(t, ps.AddCar(1))
	assert.True(t, ps.AddCar(2))
	assert.False(t, ps.AddCar(3))
	assert.False(t, ps.AddCar(1))
}

type ParkingSystem struct {
	b, m, s int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{b: big, m: medium, s: small}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.b == 0 {
			return false
		}

		ps.b--
		return true
	case 2:
		if ps.m == 0 {
			return false
		}

		ps.m--
		return true
	case 3:
		if ps.s == 0 {
			return false
		}

		ps.s--
		return true
	default:
		panic("invalid type")
	}
}
