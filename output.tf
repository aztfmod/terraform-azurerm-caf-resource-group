output "object" {
  depends_on = [azurerm_resource_group.rg]

  value = azurerm_resource_group.rg
}

output "names" {
  depends_on = [azurerm_resource_group.rg]

    value = {
    for group in keys(azurerm_resource_group.rg):
     group => azurerm_resource_group.rg[group].name
  }
}

output "ids" {
  depends_on = [azurerm_resource_group.rg]

  value = {
    for group in keys(azurerm_resource_group.rg):
     group => azurerm_resource_group.rg[group].id
}
}


