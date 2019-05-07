package acceptance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = BeforeSuite(func() {
	pathToMain, err := gexec.Build("github.com/polar-beard/terraform-diff")
	Expect(err).NotTo(HaveOccurred())
})
