package layer4

type Result int

const (
	Passed Result = iota
	Failed
	NeedsReview
	NotApplicable
	Unknown
)

func (r Result) String() string {
	// Only these three values are allowed, with all others falling back to "Unknown"
	switch r {
	case Passed:
		return "Passed"
	case Failed:
		return "Failed"
	case NeedsReview:
		return "Needs Review"
	case NotApplicable:
		return "Not Applicable"
	default:
		return "Unknown"
	}
}

// checkResultOverride compares the current result with the new result and returns the most severe of the two.
func checkResultOverride(previous Result, new Result) Result {
	if previous == Failed || new == Failed {
		// Failed should overwrite anything and immediately stop execution.
		return Failed
	}

	if previous == Unknown || new == Unknown {
		// If the current result is Unknown, it should not be overwritten by NeedsReview or Passed.
		return Unknown
	}
	if previous == NeedsReview || new == NeedsReview {
		// If the current result is NeedsReview, it should not be overwritten by Passed.
		return NeedsReview
	}
	return Passed
}
