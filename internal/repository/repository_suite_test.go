package repository_test

import (
	"testing"

	kInit "github.com/kovercjm/tool-go/init"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-auth-service/internal/repository"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var r *repository.Repository

var _ = BeforeSuite(func() {
	logger, err := kInit.DefaultLogger()
	Ω(err).ShouldNot(HaveOccurred())
	repo, err := repository.New(logger)
	Ω(err).ShouldNot(HaveOccurred())
	r = repo.(*repository.Repository)
})
