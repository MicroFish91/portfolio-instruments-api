package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

var (
	svs_token    string
	svs_testuser types.User
)

func TestSnapshotValueService(t *testing.T) {
	t.Run("Setup", snapshotValueServiceSetup)
}

func snapshotValueServiceSetup(t *testing.T) {
	svs_testuser, svs_token = NewUserSetup(t)
}
