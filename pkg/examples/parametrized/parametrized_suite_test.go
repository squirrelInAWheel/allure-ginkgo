package parametrized_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/squirrelInAWheel/allure-ginkgo/pkg/allure"
)

func TestParametrized(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parametrized Suite")
}

var _ = ReportAfterSuite("This description does nothing", func(report Report) {
	_ = allure.FromGinkgoReport(report)
})
