
resource "azurerm_resource_group" "rg" {
   for_each = var.resource_groups

         name     = "${var.prefix}${each.value.name}"
         location = each.value.location
         tags     = lookup(each.value, "tags", null) == local.tags ? null : merge(local.tags,each.value.tags)

}
