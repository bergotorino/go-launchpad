# Go Launchpad

This is a Go library for accessing [Launchpad](www.launchpad.net) suite of
tools

## Installation

```
go get github.com/bergototino/go-launchpad/launchpad
```

## Dependencies

It requires a slightly modified version of garyburd/go-oauth. You can fetch it
by typing:

```
go get github.com/bergotorino/go-oauth/oauth
```

## Level of Support

The level of support is quite limited at the moment however it will do for a
simple applications offering read only functionality over data sets such as:

* People
* Merge Proposals
* Git Repositories

## Quick Start

For a quick start check out the [Login Example](https://github.com/bergotorino/go-launchpad/blob/master/examples/login/main.go)
to learn the fundamentals. As a next step read the code of a [More Advanced Example](https://github.com/bergotorino/go-launchpad/blob/master/examples/flightschedule/main.go)
to get the feeling for overall.

## License

Go-Launchpad is available under the [GNU General Public License version 3](https://www.gnu.org/licenses/gpl-3.0.en.html)
