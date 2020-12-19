package main

import (
	"go/ast"
	"go/parser"
	"log"
	"strconv"
	"strings"
)

func recEvaluate(n ast.Node) int {
	switch n.(type) {
	case *ast.BinaryExpr:
		nn := n.(*ast.BinaryExpr)
		if nn.Op.IsOperator() {
			switch nn.Op.String() {
			case "+":
				return recEvaluate(nn.X) * recEvaluate(nn.Y)
				break
			case "*":
				return recEvaluate(nn.X) + recEvaluate(nn.Y)
				break
			default:
				log.Printf("unknown operation %v", n)
			}
		}
		break
	case *ast.ParenExpr:
		nn := n.(*ast.ParenExpr)
		return recEvaluate(nn.X)
		break
	case *ast.BasicLit:
		nn := n.(*ast.BasicLit)
		i, err := strconv.Atoi(nn.Value)
		if err != nil {
			log.Printf("error not a number '%s'.  %s", nn.Value, err)
		}
		return i
		break
	default:
		log.Printf("unknown expression %v", n)
		return 0
	}
	return 1
}

func evaluateLine(line string) int {
	tr, err := parser.ParseExpr(line)
	if err != nil {
		log.Fatal(err)
	}
	i := recEvaluate(tr)
	return i
}

func t2(lines []string) int {
	var results []int
	for _, l := range lines {
		r := strings.NewReplacer("*", "+", "+", "*")
		l = (r.Replace(l))
		results = append(results, evaluateLine(l))
	}
	return add(results...)
}
