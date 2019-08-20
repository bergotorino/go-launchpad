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

## Examples

This section documents the examples

### Manifest

The `examples/manifest` example provides a sample code that is using the
go-launchpad bindings to provide a list of downloadable URLs pointing to
archives of each source package that is used to build a selected core or
core18 snap.

Usage:
```
λ manifest (master) ✗ ./manifest --help
Usage of ./manifest:
  -revision int
          The revision to get source for (default 1076)
	    -snap string
	            The snap to get source for (default "core18")
```

Output:

```
(...)
iputils 3:20121221-5ubuntu2 https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/iputils/3:20121221-5ubuntu2/iputils_20121221.orig.tar.bz2
attr 1:2.4.47-2 https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/attr/1:2.4.47-2/attr_2.4.47.orig.tar.bz2
gmp 2:6.1.0+dfsg-2 https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/gmp/2:6.1.0+dfsg-2/gmp_6.1.0+dfsg.orig.tar.xz
jinja2 2.8-1ubuntu0.1 https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/jinja2/2.8-1ubuntu0.1/jinja2_2.8.orig.tar.gz
```

The above can be easily parsed with awk and downloaded with wget.
