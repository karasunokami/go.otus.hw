package envdir

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatalf("Error : %s", "util requires 2 args")
		return
	}

	env, err := ReadDir(args[0])
	if err != nil {
		log.Fatalf("Error : %s", err)
	}

	code := RunCmd(args[1:], env)
	os.Exit(code)
}
