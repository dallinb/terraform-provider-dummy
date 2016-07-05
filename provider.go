package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

// Provider is a basic structure that describes a provider: the configuration
// keys it takes, the resources it supports, a callback to configure, etc.
func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "dummy_server": resourceServer(),
        },
    }
}
