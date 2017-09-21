package main

import (
	"fmt"
	"os"
)

type test struct {
	Name     string `sexpr:"na"`
	Age      int
	Children []string `sexpr:"Child"`
}

func main() {
	t := test{"TestName", 18, []string{"ch1", "ch2"}}
	bytes, err := Marshal(t)
	if err != nil {
		fmt.Errorf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("encoded: %s\n", string(bytes))
	sexpr := `((na "TestName") (Age 20) (Child ("son" "daughter")))`
	var t2 test
	if err = Unmarshal([]byte(sexpr), &t2); err != nil {
		fmt.Errorf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("decoded: %v\n", t2)
}
