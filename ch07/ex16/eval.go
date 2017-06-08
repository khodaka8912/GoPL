package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calc", form)
	http.HandleFunc("/eval", eval)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const formHtml = `
<html>
<title>calculator</title>
<body>
<form method="GET" action="/eval">
<input type="text" name="expr" />
<input type="submit" value="calc" />
</form>
</body>
</html>
`

var answerHtml = template.Must(template.New("answer").Parse(`
<html>
<title>answer</title>
<body>
<p>
{{.}}
</p>
<p><a href="javascript:history.go(-1)">Back</a></p>
</body>
</html>
`))

func form(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, formHtml)
}

func eval(w http.ResponseWriter, req *http.Request) {
	expr := req.URL.Query().Get("expr")
	parsed, err := Parse(expr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad expr: %v", err)
		return
	}
	env := Env{}
	values := req.URL.Query()
	for k, v := range values {
		if k == "expr" {
			break
		}
		f, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "bad var: %s=%s", k, v)
			return
		}
		env[Var(k)] = f
	}
	vars := map[Var]bool{}
	if err = parsed.Check(vars); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad expr: %v", err)
		return
	}
	for v := range vars {
		if _, ok := env[v]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "bad var: %s", v)
			return
		}
	}
	ans := parsed.Eval(env)
	result := fmt.Sprintf("%s = %.2f", parsed, ans)
	answerHtml.Execute(w, result)
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
	return env[v]
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
