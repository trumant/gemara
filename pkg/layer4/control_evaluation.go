package layer4

import (
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"syscall"
)

// ControlEvaluation is a struct that contains all assessment results, organinzed by name
type ControlEvaluation struct {
	Name          string                      // TestSuiteName is the human-readable name or description of the control evaluation
	ControlID     string                      // ControlID is the unique identifier for the control being evaluated
	Passed        bool                        // Passed is true if all testSets in the testSuite passed
	Message       string                      // Message is the human-readable result of the final assessment to run in this evaluation
	BadStateAlert bool                        // BadState is true if any testSet failed to revert at the end of the testSuite
	Results       map[string]AssessmentResult // TestSetResults is a map of testSet names to their results
	UserGuide     string                      // UserGuide is the URL to the documentation for this evaluation
}

// ExecuteTest is a helper function to run a test function and update the result
func (c *ControlEvaluation) ExecuteTest(testFunc func() AssessmentResult) {
	// get name of the provided function as a string
	testFuncName := runtime.FuncForPC(reflect.ValueOf(testFunc).Pointer()).Name()
	// get the last part of the name, which is the actual function name
	testName := strings.Split(testFuncName, ".")[len(strings.Split(testFuncName, "."))-1]

	testResult := testFunc()

	// if this is the first test or previous tests have passed, accept any results
	if len(c.Results) == 0 || c.Passed {
		c.Passed = testResult.Passed
		c.Message = testResult.Message
	}
	c.Results[testName] = testResult
}

func (c *ControlEvaluation) Cleanup() {
	for _, result := range c.Results {
		badState := result.RevertChanges()
		if badState {
			c.BadStateAlert = true
		}
	}
}

// CloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS.
// If an interrupt is received, the program will attempt to revert any changes
// made by the terminated Plugin.
func (c *ControlEvaluation) CloseHandler() {
	// Ref: https://golangcode.com/handle-ctrl-c-exit-in-terminal/
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Print("Unexpected termination. Attempting to revert changes made by terminated Plugin. Do not interrupt this process.")
		c.Cleanup()
		os.Exit(0)
	}()
}
