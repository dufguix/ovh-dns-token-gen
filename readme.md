
# OVH api client to see and create tokens for letsencrypt DNS challenge
You should restrict credential to specific DNS API only, in case if your server is compromised.

You can follow this [Medium post](https://medium.com/nephely/configure-traefik-for-the-dns-01-challenge-with-ovh-as-dns-provider-c737670c0434) to see full instructions for Traefik.

It is quite annoying to use the official API [console](https://eu.api.ovh.com/console/) and curl command, so I created this soft.

With this soft, you can:
- See all applications infos
- See all credentials infos
- See credentials of an app
- Create credential/consumer key for DNS zone API
- Delete an app and all its credentials
- Delete a credential

## How to run
1) Create a conf file `ovh.conf` with :
```
[default]
; general configuration: default endpoint
endpoint=ovh-eu

[ovh-eu]
; configuration specific to 'ovh-eu' endpoint
application_key=YOUR_APP_KEY
application_secret=YOUR_APP_SECRET
consumer_key=YOUR_CONSUMER_KEY
```

2) Create access tokens and complete the conf file.
[OVH create a token](https://www.ovh.com/auth/api/createToken)
App name: ServerXYZ Traefik
App description: for letsencrypt DNS challenge
Validity: unlimited
Rights: get: *

Later, if you need to delete app or cred. Create with
App name: temp-client
Validity: 1day
Rigths:
- get: *
- delete: *

3) Run the soft.

## Known issues
- You may lose your rights after creating a new credential/consumer key. Just restart the script.

## Builds
You will find ovh-dns-api.exe for Windows.
For Mac and Linux, you can build with `go build .`

## Resources
[OVH API console](https://eu.api.ovh.com/console/)
[OVH go client](https://github.com/ovh/go-ovh)
[Medium post](https://medium.com/nephely/configure-traefik-for-the-dns-01-challenge-with-ovh-as-dns-provider-c737670c0434)