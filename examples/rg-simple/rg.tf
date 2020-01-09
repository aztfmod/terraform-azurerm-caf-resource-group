
module "rg-test" {
  source = "../../"
  
    prefix = local.prefix
    resource_groups = local.resource_groups
    tags = local.tags
}