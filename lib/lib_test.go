package lib

import (
	"testing"

	"github.com/mdw-go/testing/v2/should"
	"github.com/mdw-go/testing/v2/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(&Suite{T: suite.New(t)}, suite.Options.UnitTests())
}

type Suite struct {
	*suite.T
}

func (this *Suite) Setup() {}

func (this *Suite) Test() {
	this.So(1, should.Equal, 1)
}
