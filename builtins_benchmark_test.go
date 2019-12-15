package jsonnet

import (
	// "os"
	"os/exec"
	"testing"
)

func Benchmark_Builtin_substr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cmd := exec.Command("./jsonnet", "./builtin-benchmarks/substr.jsonnet")
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			b.Fail()
		}
	}
}
