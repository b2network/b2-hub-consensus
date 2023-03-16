package forkchoice

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v3/runtime/version"
	"github.com/prysmaticlabs/prysm/v3/testing/spectest/shared/common/forkchoice"
)

func TestMinimal_Deneb_Forkchoice(t *testing.T) {
	forkchoice.Run(t, "minimal", version.Deneb)
}