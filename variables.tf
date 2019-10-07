variable "resource_groups" {
  description = "(Required) Map of the resource groups to create"
}

# Example of resource_groups data structure:
# resource_groups = {
#   apim          = { 
#                   name     = "-apim-demo"
#                   location = "southeastasia" 
#                   },
#   networking    = {    
#                   name     = "-networking-demo"
#                   location = "southeastasia" 
#                   },
#   insights      = { 
#                   name     = "-insights-demo"
#                   location = "southeastasia" 
#                   },
# }

variable "prefix" {
  description = "(Optional) You can use a prefix to add to the list of resource groups you want to create"
}

variable "tags" {
  description = "tags for the deployment"
  default     = {}
}
