[![Build status](https://dev.azure.com/azure-terraform/Blueprints/_apis/build/status/modules/resource_group)](https://dev.azure.com/azure-terraform/Blueprints/_build/latest?definitionId=5)
# Creates one or multiple resource groups
Sets one or more resource groups, each of them in a specific Azure region.

Reference the module to a specific version (recommended):
```hcl
module "resource_groups" {
    source                  = "git://github.com/aztfmod/resource_group.git?ref=v1.1"
  
    prefix                  = var.prefix
    resource_groups         = var.resource_groups
    tags                    = var.tags
}
```

Or get the latest version
```hcl
module "resource_groups" {
    source                  = "git://github.com/aztfmod/resource_group.git?ref=latest"
  
    prefix                  = var.prefix
    resource_groups         = var.resource_groups
    tags                    = var.tags
}
```

# Parameters

## resource_groups
(Required) Object that describes the resource groups to be created with their geo. 
Global group of tags will be added to all RG, specific tags can be added per RG.

```hcl
variable "resource_groups" {
  description = "(Required) Map of the resource groups to create"
}
```

Example of structure: 
```hcl
resource_groups = {
    apim          = { 
                    name     = "-apim-demo"
                    location = "southeastasia" 
    },
    networking    = {    
                    name     = "-networking-demo"
                    location = "eastasia" 
    },
    insights      = { 
                    name     = "-insights-demo"
                    location = "francecentral" 
                    tags     = {
                      project     = "Pattaya"
                      approver     = "Gunter"
                    }   
    },
}
```

## prefix
(Optional) You can use a prefix to add to the list of resource groups you want to create
```hcl
variable "prefix" {
    description = "(Optional) You can use a prefix to add to the list of resource groups you want to create"
}
```
Example
```hcl
locals {
    prefix = "${random_string.prefix.result}-"
}

resource "random_string" "prefix" {
    length  = 4
    upper   = false
    special = false
}
```

# Output
## object
Returns the full set of created resource groups as a map, as follows:
```hcl
object = {
  "apim" = {
    "id" = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tst-apim-demo"
    "location" = "southeastasia"
    "name" = "tst-apim-demo"
    "tags" = {
      "BusinessUnit" = "SHARED"
      "DR" = "NON-DR-ENABLED"
      "costCenter" = "1664"
      "deploymentType" = "Terraform"
      "environment" = "DEV"
      "owner" = "Arnaud"
    }
  }
  "networking" = {
    "id" = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tst-networking-demo"
    "location" = "eastasia"
    "name" = "tst-networking-demo"
    "tags" = {
      "BusinessUnit" = "SHARED"
      "DR" = "NON-DR-ENABLED"
      "costCenter" = "1664"
      "deploymentType" = "Terraform"
      "environment" = "DEV"
      "owner" = "Arnaud"
    }
  }
}
```

## names
Kept for backward compatibility, might be removed in a future version.

Returns a map of:
- key   = key name (internal name) of the resource group
- value = name of the resource group



## ids
Kept for backward compatibility, might be removed in a future version.

Returns a map of:
- key   = key name (internal name) of the resource group
- value = id of the resource group

