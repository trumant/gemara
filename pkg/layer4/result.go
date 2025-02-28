package layer4

type Result int

const (
	Passed Result = iota
	Failed
	NeedsReview
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
	default:
		return "Unknown"
	}
}
