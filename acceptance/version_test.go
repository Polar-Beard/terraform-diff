package acceptance

import (
	"fmt"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Version", func() {

	Context("when given the -v short flag", func() {
		It("returns the binary version", func() {
			version := fmt.Sprintf("v0.0.0-dev.%d", time.Now().Unix())

			var err error
			pathToMain, err = gexec.Build("github.com/polar-beard/terraform-diff",
				"--ldflags", fmt.Sprintf("-X main.version=%s", version),
			)
			Expect(err).NotTo(HaveOccurred())

			cmd := exec.Command(pathToMain, "-v")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out).Should(gbytes.Say(version))
		})
	})

})
