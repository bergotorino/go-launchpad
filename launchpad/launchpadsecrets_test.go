package launchpad_test

import (
	"bytes"
	"encoding/json"
	. "gopkg.in/check.v1"

	"github.com/bergotorino/go-launchpad/launchpad"
	"github.com/bergotorino/go-oauth/oauth"
)

type LaunchpadSecretsSuite struct {
	ac   oauth.Credentials
	data []byte
}

var _ = Suite(&LaunchpadSecretsSuite{})

func (s *LaunchpadSecretsSuite) SetUpTest(c *C) {
	s.ac.Token = "Gl70PLC97z0TG6JdMCKs"
	s.ac.Secret = "5GD8s3TBxQvNXSLML0K7HgzRwHlKCC9LTXN00nRW2HzKc2DBKS"
	var err error
	s.data, err = json.Marshal(s.ac)
	c.Assert(err, IsNil)
}

func (s *LaunchpadSecretsSuite) TestReadAndWrite(c *C) {
	var ls launchpad.LaunchpadSecrets
	var err error

	// Test reading
	rbuffer := bytes.NewBuffer(s.data)
	err = ls.Read(rbuffer)
	c.Assert(err, IsNil)
	c.Assert(ls.IsValid(), Equals, true)

	// Test writing
	var wbuffer bytes.Buffer
	err = ls.Write(&wbuffer)
	c.Assert(err, IsNil)
	c.Assert(wbuffer.Bytes(), DeepEquals, s.data)
}
