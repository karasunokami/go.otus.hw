package envdir

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func ReadDir(dir string) (map[string]string, error) {
	env := make(map[string]string)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return env, err
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return env, err
		}
		if string(content) == "" {
			continue
		}
		env[file.Name()] = string(content)
	}
	return env, err
}

func RunCmd(cmd []string, env map[string]string) int {
	command := exec.Command(cmd[0])
	if len(cmd) > 1 {
		command.Args = cmd[1:]
	}
	result := make([]string, 0, len(env))
	for key, value := range env {
		result = append(result, key+"="+value)
	}
	command.Env = append(os.Environ(), result...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}

	return 0
}
