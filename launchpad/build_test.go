package launchpad_test

import (
	. "gopkg.in/check.v1"

	"github.com/bergotorino/go-launchpad/launchpad"
)

type BuildSuite struct {
	build launchpad.Build
}

var _ = Suite(&BuildSuite{})

func (s *BuildSuite) TestBuildState(c *C) {
	var buildTests = []struct {
		value    string
		expected int
	}{
		{"Failed to build", launchpad.BuildFailure},
		{"Successfully built", launchpad.BuildSuccess},
		{"Needs building", launchpad.BuildNeedsBuilding},
		{"Dependency wait", launchpad.BuildDependencyWait},
		{"Chroot problem", launchpad.BuildChrootProblem},
		{"Build for superseded Source", launchpad.BuildBuildSupersededSource},
		{"Currently building", launchpad.BuildBuilding},
		{"Failed to upload", launchpad.BuildUploadFailed},
		{"Uploading build", launchpad.BuildUploading},
		{"Cancelling build", launchpad.BuildCancelling},
		{"Cancelled build", launchpad.BuildCancelled},
		{"Foo Bar", launchpad.BuildUnknown},
	}

	for _, bt := range buildTests {
		s.build.Buildstate = bt.value
		c.Assert(s.build.BuildState(), Equals, bt.expected)
	}
}
