# proxmoxve-client-go

WIP Go client for ProxMox Virtual Environment's JSON API.

This will be used in a Terraform provider so from this point on I will prioritise those APIs that I need. Contributions are welcome.

Instructions for running tests will be added in the future.

## Testing

There's a little set up needed before the test suite can be run.

### Proxmox server

You need a Proxmox VE server running somewhere. I suggest a throwaway VM since we'll be creating and deleting resources, adding a trusted CA and other things that shouldn't ever be done in a production system.

Create an API token for the `root@pam` user.

Export the following environment variables:

```bash
export PROXMOXVE_TEST_BASE_URL="https://<pmve_host>:<port>"
export PROXMOXVE_TEST_TOKEN_ID="<root_token_id>"
export PROXMOXVE_TEST_SECRET="<root_token_secret>"
export PROXMOXVE_TEST_TLS_INSECURE=true

# This is a safeguard. It's used my `make cleanup` and means you are aware it might delete things it shouldn't. Please run this against a throwaway VM.
export PROXMOXVE_TEST_URL_CLEANUP="https://<pmve_host>:<port>"

# Some endpoints require this, e.g. ACME account creation. See https://forum.proxmox.com/threads/acme-api-endpoint-403-permission-check-failed-user-root-pam-despite-user-being-root-pam.111745/
export PROXMOXVE_TEST_USER="root@pam"
export PROXMOXVE_TEST_PASS="<root_password>"
export PROXMOXVE_TEST_TOTPSEED="<root_totp_seed>"
```

### Run Pebble

Install and run [Let's Encrypt Pebble](https://github.com/letsencrypt/pebble) on the host.

```bash
apt-get update
apt-get -y install golang git
git clone https://github.com/letsencrypt/pebble.git
cd pebble
go install ./cmd/pebble
cp ~/pebble/test/certs/pebble.minica.pem /usr/local/share/ca-certificates/pebble.minica.crt
update-ca-certificates --fresh
~/go/bin/pebble -config ./test/config/pebble-config.json
```

Then, create the ACME account:

```
pvesh create /cluster/acme/account --name "pebble" --contact "foo@bar.com" --directory "https://localhost:14000/dir" --tos_url "foobar"
```
