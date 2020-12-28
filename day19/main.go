package main

import (
	"fmt"
)

const (
	condition = iota
	basicexpr
)

type Kind int

type Facts map[int]bool

func (all Facts) check(f int) bool {
	if r, ok := all[f]; ok {
		return r
	} else {
		return false
	}
}

type Rule struct {
	number        int
	preconditions [][]int
	literal       rune
	kind          Kind
}

func NewRule(number int, checks [][]int, literal rune) *Rule {
	r := &Rule{number: number, preconditions: checks, literal: literal}
	if checks != nil {
		r.kind = condition
	} else {
		r.kind = basicexpr
	}
	return r
}

func checkpreconditions(f Facts, preconditions []int) bool {
	//or condition
	for _, p := range preconditions {
		if f.check(p) {
			return true
		}
	}
	return false
}

func hypothesizeAndTest (startnodes rules preconditions) {

}

/*
(defun hypothesizse-and-test (startnodes rules preconditions)
  (let ((facts'())
        (hypotheses  (sort startnodes #'string<)))
    (do* ((i 0 (1+ i))
          (currenthypothesis (nth i (sort hypotheses #'string<)) (nth i (sort hypotheses #'string<)))
          (currentrules (gethash currenthypothesis rules)  (gethash currenthypothesis rules)))
         ((>= i (length hypotheses)))
      (when (checkpreconditions facts (gethash currenthypothesis preconditions))
        (setq hypotheses (append (subseq hypotheses 0 i) (subseq hypotheses (+ 1 i))))
        (setq i -1)
        (setq hypotheses (sort (union currentrules hypotheses :test #'string=) #'string<))
        (setq facts (cons currenthypothesis facts))
        ))
    (format NIL "~{~a~}" (reverse facts))
    ))
*/

func t1() {

}

func main() {
	var input string = "input.txt"
	rules, facts := ParseData(ReadData(input))

	fmt.Printf("ticket scanning error rate: %d\n", t1(rules, myticket, nearbies))
	//fmt.Printf("Prouct of the six departure numbers: %d\n", ProductOfDepartureField(rules, myticket, nearbies))
}
