package snapshot

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/snapshot/advanced"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/snapshot/core"
)

func TestSnapshotService(t *testing.T) {
	t.Run("Scenario: Core", core.CoreSnapshotScenarioTests)
	t.Run("Scenario: Advanced", advanced.AdvancedSnapshotScenarioTests)
}
