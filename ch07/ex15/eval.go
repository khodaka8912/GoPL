package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		parsed, err := Parse(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v", err)
			continue
		}
		env := Env{}
		ans := parsed.Eval(env)
		fmt.Printf("%s = %.2f %v\n", parsed, ans, env)
	}
}

func (e Env) String() string {
	buf := bytes.NewBufferString("[")
	for k, v := range e {
		if buf.Len() > 1 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s=%.2f", k, v)
	}
	buf.WriteRune(']')
	return buf.String()
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (v Var) String() string {
	return string(v)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteByte(')')
	return buf.String()
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	val, ok := env[v]
	if ok {
		return val
	}
	fmt.Printf("Input value of %q: ", v)
	_, err := fmt.Scanf("%f", &val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	env[v] = val
	return val
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error

	String() string
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
