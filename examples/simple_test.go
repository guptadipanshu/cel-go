package examples

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func ExampleSimple() {
	d := cel.Declarations(decls.NewIdent("name", decls.String, nil))
	env, err := cel.NewEnv(d)
	if err != nil {
		log.Fatalf("environment creation error: %v\n", err)
	}
	// Check iss for error in both Parse and Check.
	ast, iss := env.Compile(`"Hello world! I'm " + name + "."`)
	if iss != nil && iss.Err() != nil {
		log.Fatalln(iss.Err())
	}
	prg, err := env.Program(ast)

	out, _, err := prg.Eval(map[string]interface{}{
		"name": "CEL",
	})
	fmt.Println(out)
	// Output:Hello world! I'm CEL.
}
