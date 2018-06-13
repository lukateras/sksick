Thanks a whole lot to Micah Lee who has tweeted about this bug in the first place.

# SKS Tools

Tools to break SKS PGP keyserver sites. It is possible to:

- Cause denial of service (anyone can make `gpg --recv-key` fail for any key):
  https://bitbucket.org/skskeyserver/sks-keyserver/issues/57/anyone-can-make-any-pgp-key-unimportable
- Cause major rendering bugs on some PGP keyserver sites (at least
  pgp.key-server.io) that make it impossible to read any information
  on the page
- Add a fake UID to any PGP key


