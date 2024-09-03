package snapshot

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/snapshot/advanced"
)

func TestSnapshotService(t *testing.T) {
	// Basic

	// Scenarios
	t.Run("Scenario: Advanced", advanced.AdvancedSnapshotScenarioTests)
}
