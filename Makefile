default: sks-fake-uid/sks-fake-uid

src:
	ln -s . src

sks-fake-uid/sks-fake-uid: sks-fake-uid/*.go src
	cd sks-fake-uid; GOPATH=$(shell pwd) go build
