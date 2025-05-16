package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTask4(t *testing.T) {
    // Read the test input file
    testFile := "tests/1"
    inputData, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatalf("Failed to read test file: %v", err)
    }

    // Create input and output buffers
    inputBuffer := strings.NewReader(string(inputData))
    outputBuffer := &bytes.Buffer{}
    
    // Run the solution with the test input
    runSolution(inputBuffer, outputBuffer)
    
    // Get the output
    output := outputBuffer.String()
    
    // Here you can check if the output matches the expected output
    t.Logf("Output: %s", output)
    
    // Example assertion (replace with actual expected output)
    // expected := "NO\nNO\nNO\n"
    // if output != expected {
    //     t.Errorf("Expected output %q but got %q", expected, output)
    // }
}