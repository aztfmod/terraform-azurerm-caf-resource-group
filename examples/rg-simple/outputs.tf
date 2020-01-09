output "object" {
    description = "Returns the full set of resource group objects created"

    value = module.rg-test.object
}

output "names" {
    description = "Returns a map of resource_group key -> resource_group name"

    value = module.rg-test.names
}

output "ids" {
    description = "Returns a map of resource_group key -> resource_group id"

    value = module.rg-test.ids
}


