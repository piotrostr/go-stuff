package website_racer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWebsiteRacer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WebsiteRacer Suite")
}
