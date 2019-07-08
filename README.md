## Install

```
go build
```

## Run

```
# Get provider (need to run after a `go build`):
terraform init

# Apply state:
terraform apply -auto-approve

# Destroy project:
terraform destroy
```

## Limitations

* work only with terraform <= 0.11
* build with golang 1.12.6 (may be work with another versions)
* delete work only on project
* no update on state allowed
* A problem with environnement (need to run 1 time, clean manualy state, run another time)
