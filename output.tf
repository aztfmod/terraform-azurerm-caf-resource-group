output "object" {
  description = "Returns the full set of resource group objects created"
  depends_on = [azurerm_resource_group.rg]

  value = azurerm_resource_group.rg
}

output "names" {
  description = "Returns a map of resource_group key -> resource_group name"
  depends_on = [azurerm_resource_group.rg]

    value = {
    for group in keys(azurerm_resource_group.rg):
     group => azurerm_resource_group.rg[group].name
  }
}

output "ids" {
  description = "Returns a map of resource_group key -> resource_group id"
  depends_on = [azurerm_resource_group.rg]

  value = {
    for group in keys(azurerm_resource_group.rg):
     group => azurerm_resource_group.rg[group].id
}
}


