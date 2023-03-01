package judge

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func JudgeCode(sourceCodeFilePath string, testCaseFilePath string) string {
	pwd, _ := os.Getwd()
	println(pwd + testCaseFilePath)
	fi, err := os.Open(pwd + testCaseFilePath)
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	txt := ""
	out := ""

	for scanner.Scan() {
		tmp := scanner.Text()
		if tmp == "" {
			continue
		}
		if tmp == "#END_OF_IN" {
			out = runCode(txt, pwd+sourceCodeFilePath)
			txt = ""
		} else if tmp == "#END_OF_OUT" {
			if txt != out {
				fmt.Printf("Expected output:\n%s\n", txt)
				fmt.Printf("Wrong Answer\n")
				return "Wrong Answer"

			}
			txt = ""
		} else {
			txt += tmp + "\n"
		}
	}

	fmt.Printf("Accepted\n")

	return "Accepted"

}

func runCode(input string, sourceCodeFilePath string) string {
	pwd, _ := os.Getwd()
	cmd := exec.Command(pwd+"/judge/test.sh", sourceCodeFilePath, input)
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
