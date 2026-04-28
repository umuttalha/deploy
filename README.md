# deploy

CLI scaffold for an infra provisioner that uses the AWS and Cloudflare SDKs directly. State is tracked via resource tags (`managed-by=deploy`, `stack=<name>`). **Boilerplate only — no resource is actually provisioned yet.**

## Build

```bash
make build
./bin/deploy --help
```

## Commands

```bash
deploy up [stack-name]    # provision (interactive prompts for missing options)
deploy down <stack-name>  # tear down
deploy ls                 # list every stack tagged managed-by=deploy
deploy version
```

Aliases: `deploy` → `up`, `destroy` → `down`, `list` → `ls`.

## Test

```bash
make test
```
