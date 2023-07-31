package lib

import (
	"testing"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {}

func (this *Suite) Test() {
	this.So(must.Value("hello", nil), should.Equal, "hello")
	this.So(funcy.Sum([]int{1, 2, 3}), should.Equal, 6)
	this.So(set.Of[int](1, 1, 2, 2, 3, 3).Len(), should.Equal, 3)
}
