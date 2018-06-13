package main

import (
	"flag"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"log"
	"math/rand"
	"os"
)

var randOn = flag.Bool("rand", false, "use random bytes instead of UID name, comment and email")
var randLen = flag.Uint64("randLen", 100000, "length of each random UID fragment")
var uidName = flag.String("uidName", "Do not use SKS keyserver sites", "UID name")
var uidComment = flag.String("uidComment", "no validity checks", "UID comment")
var uidEmail = flag.String("uidEmail", "https://bitbucket.org/skskeyserver/sks-keyserver/issues/41", "UID email")

func randString(length uint64) string {
	buf := make([]byte, length)
	rand.Read(buf) // doesn't really fail
	return string(buf)
}

func main() {
	flag.Parse()

	entity, err := openpgp.ReadEntity(packet.NewReader(os.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	var sig *packet.Signature

	for _, v := range entity.Identities {
		sig = v.SelfSignature
	}

	var pak *packet.UserId

	if *randOn {
		pak = packet.NewUserId(
			randString(*randLen),
			randString(*randLen),
			randString(*randLen))
	} else {
		pak = packet.NewUserId(*uidName, *uidComment, *uidEmail)
	}

	entity.Identities[pak.Id] = &openpgp.Identity{
		Name:          pak.Id,
		UserId:        pak,
		SelfSignature: sig,
	}

	for _, v := range entity.Identities {
		entity.Identities[pak.Id].Signatures =
			append(entity.Identities[pak.Id].Signatures, v.Signatures...)
	}

	err = entity.Serialize(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
