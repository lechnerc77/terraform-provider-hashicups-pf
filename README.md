# Terraform Provider Hashicups

This repo is a companion repo to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers) Learn collection. 

In the collection, you will use the HashiCups provider as a bridge between Terraform and the HashiCups API. Then, extend Terraform by recreating the HashiCups provider. By the end of this collection, you will be able to take these intuitions to create your own custom Terraform provider. 

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-hashicups
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

## Setup local provider

Place the `terraform.rc` file into `C:\Users\I550309\AppData\Roaming`

## API backend

Check <https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider-configure#start-hashicups-locally> for instruction on how to start API backend
