package integration_test

import (
	"github.com/cloudfoundry/libbuildpack/cutlass"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path/filepath"
)

var _ = Describe("setting url/find links in requirements.txt", func() {
	var app *cutlass.App

	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	FContext("find-links is set", func() {
		BeforeEach(func() {
			app = cutlass.New(filepath.Join(bpDir, "fixtures", "vendored_transitive_dependencies"))
			app.SetEnv("BP_DEBUG", "1")
		})

		It("should work", func() {
			PushAppAndConfirm(app)
			Expect(app.GetBody("/")).To(ContainSubstring("Welcome to Python on Cloud Foundry"))
		})

		AssertNoInternetTraffic("vendored_transitive_dependencies")
	})

	Context("index-url is set", func() {

	})

	Context("extra-index-url is set", func() {

	})

	//Context("buildpack is cached", func() {
	//	Context("python 3", func() {
	//		Context("app is completely vendored", func() {
	//			BeforeEach(func() {
	//				if !cutlass.Cached {
	//					Skip("Running cached tests")
	//				}
	//				app = cutlass.New(filepath.Join(bpDir, "fixtures", "flask_python_3_pipenv_vendored"))
	//				app.SetEnv("BP_DEBUG", "1")
	//			})
	//
	//			It("should work", func() {
	//				PushAppAndConfirm(app)
	//				Expect(app.GetBody("/")).To(ContainSubstring("Hello, World with pipenv!"))
	//			})
	//
	//			AssertNoInternetTraffic("flask_python_3_pipenv_vendored")
	//		})
	//	})
	//})
})
