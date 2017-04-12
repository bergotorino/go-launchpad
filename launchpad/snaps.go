package launchpad

import (
	"log"
	"net/url"
)

type Snaps struct {
	lp *Launchpad
}

func (s *Snaps) GetByName(name string, owner string) (*Snap, error) {
	v := url.Values{}
	v.Set("name", name)
	v.Add("owner", owner)
	v.Add("ws.op", "getByName")

	response, err := s.lp.Get("https://api.launchpad.net/devel/+snaps", v)
	if err != nil {
		log.Println("API returned failure", err)
		return nil, err
	}

	var data Snap
	err = DecodeResponse(response, &data)
	if err != nil {
		log.Println("Decoding went bad")
		return nil, err
	}

	data.lp = s.lp

	return &data, nil
}
