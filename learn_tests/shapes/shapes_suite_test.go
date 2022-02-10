package shapes_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShapes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shapes Suite")
}
