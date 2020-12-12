package main

import (
	"fmt"
	"regexp"
	"strings"

	"log"
)

var regx1 = regexp.MustCompile(`^([NESWLRF]+)(\d+)$`)

type Command struct {
	cmd   string
	param int
}

func NewCommand(cmd string, param int) *Command {
	return &Command{cmd: cmd, param: param}
}
func (c *Command) String() string {
	return fmt.Sprintf("%s%d", c.cmd, c.param)
}

func callMoveCommand(cmdfs map[string]MoveFunc, c *Command, x, y, d int) (int, int, int) {
	if f, ok := cmdfs[c.cmd]; ok {
		d, x, y = f(x, y, d, c.param)
	} else {
		log.Printf("Warning: unhandled command: %s", c)
	}
	return d, x, y
}

func callWaypointCulculationCommand(wpcfs map[string]WaypointCalculationFunc, c *Command, x, y, wx, wy int) (int, int, int, int) {
	if f, ok := wpcfs[c.cmd]; ok {
		x, y, wx, wy = f(x, y, wx, wy, c.param)
	} else {
		log.Printf("Warning: unhandled command: %s", c)
	}
	return x, y, wx, wy
}
func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func travel(cmds []*Command) (x, y, r int) {
	cmdf := make(map[string]MoveFunc)
	cmdf["N"] = MoveNorth
	cmdf["E"] = MoveEast
	cmdf["S"] = MoveSouth
	cmdf["W"] = MoveWest
	cmdf["L"] = TurnLeft
	cmdf["R"] = TurnRight
	cmdf["F"] = MoveForward
	d := 90
	for _, c := range cmds {
		d, x, y = callMoveCommand(cmdf, c, x, y, d)
	}
	r = abs(x) + abs(y)
	return x, y, r
}

func waypointing(cmds []*Command) (x, y, r int) {
	movf := make(map[string]MoveFunc)
	movf["N"] = MoveNorth
	movf["E"] = MoveEast
	movf["S"] = MoveSouth
	movf["W"] = MoveWest

	wpcf := make(map[string]WaypointCalculationFunc)
	wpcf["L"] = RotateWaypointLeft
	wpcf["R"] = RotateWaypointRight
	wpcf["F"] = MoveForwardByWaypoint

	d := 90
	wx := 10
	wy := -1
	for _, c := range cmds {
		switch {
		case strings.Contains("NESW", c.cmd):
			d, wx, wy = callMoveCommand(movf, c, wx, wy, d)
			break
		case strings.Contains("LRF", c.cmd):
			x, y, wx, wy = callWaypointCulculationCommand(wpcf, c, x, y, wx, wy)
			break
		default:
			log.Printf("Warning: unhandled command in 'waypointing': %s", c)
		}
	}
	r = abs(x) + abs(y)
	return x, y, r
}

func main() {
	var input string = "input.txt"
	cmds := parsedata(readdata(input))

	x, y, r := travel(cmds)
	fmt.Printf("Simple travelling leads to: [%6d,%6d]   Manhattan distance: %6d\n", x, y, r)

	x, y, r = waypointing(cmds)
	fmt.Printf("Waypointing leads to:       [%6d,%6d]   Manhattan distance: %6d\n", x, y, r)
}
