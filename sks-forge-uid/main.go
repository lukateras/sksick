package main

import (
	"flag"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"log"
	"os"
)

func main() {
	comment := flag.String("comment", "no validity checks", "UID comment")
	email := flag.String("email", "@", "UID email")
	name := flag.String("name", "Do not use SKS keyserver sites", "UID name")

	flag.Parse()

	entity, err := openpgp.ReadEntity(packet.NewReader(os.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	var sig *packet.Signature

	for _, v := range entity.Identities {
		sig = v.SelfSignature
	}

	packet := packet.NewUserId(*name, *comment, *email)

	entity.Identities[packet.Id] = &openpgp.Identity{
		Name:          packet.Id,
		UserId:        packet,
		SelfSignature: sig,
	}

	err = entity.Serialize(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
