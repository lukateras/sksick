![SKSick](logo.png)

Tools to break SKS PGP keyserver sites. It is possible to:

- [Cause denial of service](https://bitbucket.org/skskeyserver/sks-keyserver/issues/57):
  anyone can make `gpg --recv-key` fail for any key!
- [Add fake UIDs to any PGP key](screenshots/01.png)

Thanks a whole lot to Micah Lee who has tweeted about this bug in the first place.
