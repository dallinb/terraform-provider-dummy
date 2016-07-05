# terraform-provider-dummy

## Overview
This is an implementation of
https://www.hashicorp.com/blog/terraform-custom-providers.html
and attempts to help work around a couple of gotcha's.

## Working Versions
```bash
$ terraform version
Terraform v0.6.16

$ go version
go version go1.6 linux/amd64
```

## Using This Example

Assuming that you have already got ```GOPATH``` setup (see
https://golang.org/doc/code.html for details).  Do the following:

```bash
$ go get -u github.com/hashicorp/terraform
$ go get -u github.com/dallinb/terraform-provider-dummy
```

Now get arount a problem with a new Terraform API version introduced
after 0.7.0:

```bash
$ cd $GOPATH/src/github.com/hashicorp/terraform
$ git checkout v0.6.16
Note: checking out 'v0.6.16'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by performing another checkout.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -b with the checkout command again. Example:

  git checkout -b <new-branch-name>

HEAD is now at 6e586c8... v0.6.16
```

Now let's build the plugin and try this out:

```bash
$ cd $GOPATH/src/github.com/dallinb/terraform-provider-dummy
$ go build -o terraform-provider-dummy
$ terraform plan
Refreshing Terraform state prior to plan...


The Terraform execution plan has been generated and is shown below.
Resources are shown in alphabetical order for quick scanning. Green resources
will be created (or destroyed and then created if an existing resource
exists), yellow resources are being changed in-place, and red resources
will be destroyed.

Note: You didn't specify an "-out" parameter to save this plan, so when
"apply" is called, Terraform can't guarantee this is what will execute.

+ dummy_server.foo
    address: "" => "hashicorp.com"


Plan: 1 to add, 0 to change, 0 to destroy.
```
