default: sks-fake-uid/sks-fake-uid

sks-fake-uid/sks-fake-uid: sks-fake-uid/main.go
	cd sks-fake-uid; go build
