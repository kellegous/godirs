# GoDirs

wrappers around standard [go](http://golang.org/) command line tools
that make it possible to set the GOPATH for particular working directories.
This makes it easier to manage go projects with local workspaces.

## Installing

    go get \
      github.com/kellegous/godirs/gaan \
      github.com/kellegous/godirs/gaandoc

That will create bin/gaan and bin/gaandoc in your GOPATH. These tools are simple wrappers that invoke go and godoc, respectively. The environmental
variable GOPATH is determined by looking for a .gaan file in the current
working directory or any of its parent directories.

## Example use
// TODO(kellegous): Make it happen.