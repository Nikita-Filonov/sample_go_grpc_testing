package common

import (
	"fmt"
	"github.com/onsi/gomega"
	"go.uber.org/dig"
	"sample_go_grpc_testing/utils/common/container"
	"testing"
)

func SetupTesting(t *testing.T) (*container.Components, gomega.Gomega) {
	g := GetGomega(t)

	c, err := container.BuildContainer()
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), fmt.Sprintf("unable to build container: %v", dig.RootCause(err)))

	return c, g
}
