package layer4

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ControlEvaluation is a struct that contains all assessment results, organinzed by name
type ControlEvaluation struct {
	Name            string       // TestSuiteName is the human-readable name or description of the control evaluation
	Control_Id      string       // Control_Id is the unique identifier for the control being evaluated
	Result          Result       // Result is true if all testSets in the testSuite passed
	Message         string       // Message is the human-readable result of the final assessment to run in this evaluation
	Corrupted_State bool         // BadState is true if any testSet failed to revert at the end of the testSuite
	User_Guide      string       // User_Guide is the URL to the documentation for this evaluation
	Assessments     []Assessment // Control_Evaluations is a map of testSet names to their results
}

// Evaluate runs each step in each assessment, updating the relevant fields on the control evaluation.
// It will halt if a step returns a failed result.
func (c *ControlEvaluation) Evaluate(targetData interface{}) {
	for _, assessment := range c.Assessments {
		result := assessment.Run(targetData)
		c.Result = checkResultOverride(c.Result, result)
		if c.Result == Failed {
			break
		}
	}
}

// TolerantEvaluate runs each step in each assessment, updating the relevant fields on the control evaluation
// It will not halt if a step returns an failed result.
func (c *ControlEvaluation) TolerantEvaluate(targetData interface{}) {
	for _, assessment := range c.Assessments {
		result := assessment.RunTolerateFailures(targetData)
		c.Result = checkResultOverride(c.Result, result)
	}
}

func (c *ControlEvaluation) Cleanup() {
	for _, assessment := range c.Assessments {
		corrupted := assessment.RevertChanges()
		if corrupted {
			c.Corrupted_State = true
		}
	}
}

// CloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the operating system.
// If an interrupt is received, this will attempt to revert any changes
// made by the terminated ControlEvaluation.
func (c *ControlEvaluation) CloseHandler() {
	// Ref: https://golangcode.com/handle-ctrl-c-exit-in-terminal/
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Print("\n*****\nUnexpected termination. Attempting to revert changes made by the active ControlEvaluation. Do not interrupt this process.\n*****\n")
		c.Cleanup()
		os.Exit(0)
	}()
}
