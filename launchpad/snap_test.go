package launchpad_test

import (
	. "gopkg.in/check.v1"

	"github.com/bergotorino/go-launchpad/launchpad"
)

type SnapSuite struct {
	snap launchpad.Snap
}

var _ = Suite(&SnapSuite{})

func (s *SnapSuite) SetUpTest(c *C) {
	s.snap = launchpad.Snap{WebLink: "http://foo/bar"}
}

func (s *SnapSuite) TestString(c *C) {
	c.Assert(s.snap.String(), Equals, "http://foo/bar")
}

func (s *SnapSuite) TestProcessors(c *C) {

}
