package judge

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func JudgeCode(sourceCodeFilePath string, testCaseFilePath string) {
	fi, err := os.Open(testCaseFilePath)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	txt := ""
	out := ""
	status := true

	for scanner.Scan() {
		tmp := scanner.Text()
		if tmp == "" {
			continue
		}
		if tmp == "#END_OF_IN" {
			out = runCode(txt, sourceCodeFilePath)
			txt = ""
		} else if tmp == "#END_OF_OUT" {
			if txt != out {
				status = false
				fmt.Printf("Expected output:\n%s\n", txt)
				fmt.Printf("Wrong Answer\n")
				break
			}
			txt = ""
		} else {
			txt += tmp + "\n"
		}
	}
	if status {
		fmt.Printf("Accepted\n")
	}
}

func runCode(input string, sourceCodeFilePath string) string {
	cmd := exec.Command("./test.sh", sourceCodeFilePath, input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Runtime Error: %s\n", string(out))
		fmt.Printf("Runtime Error: %s\n", string(err.Error()))
		os.Exit(0)
	}
	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Outout:\n%s\n", string(out))
	return string(out)
}
