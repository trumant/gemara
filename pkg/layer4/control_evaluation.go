package layer4

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ControlEvaluation is a struct that contains all assessment results, organinzed by name
type ControlEvaluation struct {
	Name          string                      // TestSuiteName is the name of the TestSuite
	Passed        bool                        // Passed is true if all testSets in the testSuite passed
	BadStateAlert bool                        // BadState is true if any testSet failed to revert at the end of the testSuite
	Results       map[string]AssessmentResult // TestSetResults is a map of testSet names to their results
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
