package go_utils_test

import (
	"testing"

	base "github.com/savannahghi/go_utils"
)

func TestModelsIsEntity(t *testing.T) {

	t11 := base.EmailOptIn{}
	t11.IsEntity()
}
