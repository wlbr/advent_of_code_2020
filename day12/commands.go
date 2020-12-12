package main

import (
	"log"
)

type MoveFunc func(x, y int, d int, p int) (int, int, int)

func MoveNorth(x, y int, d int, p int) (int, int, int) {
	return d, x, y - p
}

func MoveEast(x, y int, d int, p int) (int, int, int) {
	return d, x + p, y
}

func MoveSouth(x, y int, d int, p int) (int, int, int) {
	return d, x, y + p
}

func MoveWest(x, y int, d int, p int) (int, int, int) {
	return d, x - p, y
}

func TurnLeft(x, y int, d int, p int) (int, int, int) {
	nd := (d - p) % 360
	if nd < 0 {
		nd += 360
	}
	return nd, x, y
}

func TurnRight(x, y int, d int, p int) (int, int, int) {
	return abs((d + p)) % 360, x, y
}

func MoveForward(x, y int, d int, p int) (int, int, int) {
	switch d / 90 {
	case 0:
		return MoveNorth(x, y, d, p)
	case 1:
		return MoveEast(x, y, d, p)
	case 2:
		return MoveSouth(x, y, d, p)
	case 3:
		return MoveWest(x, y, d, p)
	default:
		log.Fatal("Not implemented")
		return 0, 0, 0
	}
}

type WaypointCalculationFunc func(x, y, wx, wy, p int) (int, int, int, int)

func RotateWaypointLeft(x, y, wx, wy, p int) (int, int, int, int) {
	for i := 0; i < p/90; i++ {
		dx := wx - x
		dy := wy - y
		wx = dy + x
		wy = dx*-1 + y
	}
	return x, y, wx, wy
}

func RotateWaypointRight(x, y, wx, wy, p int) (int, int, int, int) {
	for i := 0; i < p/90; i++ {
		dx := wx - x
		dy := wy - y
		wx = dy*-1 + x
		wy = dx + y
	}
	return x, y, wx, wy
}

func MoveForwardByWaypoint(x, y, wx, wy, p int) (int, int, int, int) {
	dx := wx - x
	dy := wy - y

	x = x + dx*p
	y = y + dy*p

	wx = x + dx
	wy = y + dy
	return x, y, wx, wy
}
