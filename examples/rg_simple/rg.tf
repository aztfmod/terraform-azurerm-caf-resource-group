provider "azurerm" {
  version = "=2.0.0"
  features {}
}
module "rg_test" {
  source = "../../"
  
    prefix = local.prefix
    resource_groups = local.resource_groups
    tags = local.tags
}
