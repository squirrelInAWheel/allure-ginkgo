package statuses_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/squirrelInAWheel/allure-ginkgo/pkg/allure"
)

func TestStatuses(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Statuses Suite")
}

var _ = ReportAfterSuite("This description does nothing", func(report Report) {
	_ = allure.FromGinkgoReport(report)
})
