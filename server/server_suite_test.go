package server_test

import (
	"github.com/oschwald/geoip2-golang"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *geoip2.Reader

func TestGeo2ip(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geo2ip Suite")
}

var _ = BeforeSuite(func() {
	var err error
	db, err = geoip2.Open("../data/GeoIP2-City/GeoIP2-City.mmdb")
	Expect(err).NotTo(HaveOccurred())
})
