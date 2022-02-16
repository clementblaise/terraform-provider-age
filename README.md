# Age Terraform Provider

This provider lets you generate an Age key pair.

## Using the provider

View the documentation on [Terraform registry](https://registry.terraform.io/providers/clementblaise/age/latest/docs) or in the [docs folder](docs)

```hcl
terraform {
  required_providers {
    age = {
      source = "clementblaise/age"
    }
  }
}

resource "age_secret_key" "example" {}
```

## Contributing

If you wish to work on the provider, you'll need:

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0.x
- [Go](https://golang.org/doc/install) >= 1.15

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

Add the local build to your local Terraform plugins so you can test it in your project context.

```sh
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/clementblaise/age/0.0.1/$(go env GOOS)_$(go env GOARCH)
ln -s $(go env GOPATH)/bin/terraform-provider-age ~/.terraform.d/plugins/registry.terraform.io/clementblaise/age/0.0.1/$(go env GOOS)_$(go env GOARCH)/terraform-provider-age
```

In order to run the full suite of Acceptance tests, run `make test`.

To generate or update documentation, run `make doc`.

## Maintaining

To release, just push a new tag respecting Semver:

```sh
VERSION=v0.1.0
git tag $VERSION main
git push origin $VERSION
```