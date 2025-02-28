package layer4

import "testing"

var changesTestData = []struct {
	testName string
	change   *Change
}{
	{
		testName: "Change not yet applied",
		change:   pendingChange,
	},
	{
		testName: "Change already applied and not yet reverted",
		change:   appliedNotRevertedChange,
	},
	{
		testName: "Change already applied and reverted",
		change:   appliedRevertedChange,
	},
	{
		testName: "No revert function specified",
		change:   noRevertChange,
	},
	{
		testName: "No apply function specified",
		change:   noApplyChange,
	},
	{
		testName: "Neither function specified (1)",
		change:   &Change{},
	},
}

func TestApply(t *testing.T) {
	for _, c := range changesTestData {
		t.Run(c.testName, func(t *testing.T) {

			c.change.Apply()

			if c.change.applyFunc == nil && c.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if c.change.revertFunc == nil && c.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if c.change.applyFunc != nil && c.change.revertFunc != nil {
				if !c.change.Applied {
					t.Errorf("Expected change to be applied, but it was not")
				}

				c.change.Revert()

				if !c.change.Applied {
					t.Errorf("Reverting shound not erase the record that the change was applied")
				}
			}
		})
	}
}

func TestRevert(t *testing.T) {
	for _, c := range changesTestData {
		t.Run(c.testName, func(t *testing.T) {

			c.change.Revert()

			if c.change.applyFunc == nil && c.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if c.change.revertFunc == nil && c.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if c.change.applyFunc != nil && c.change.revertFunc != nil {
				if c.change.Applied && !c.change.Reverted {
					t.Errorf("Expected change to be reverted, but it was not")
				}
				if !c.change.Applied && c.change.Reverted {
					t.Errorf("Reverting should not be recorded if a change was not applied to revert")
				}
				c.change.Apply()

				if c.change.Reverted {
					t.Errorf("Applying further times shound mark the change as not reverted")
				}
			}

		})
	}
}
