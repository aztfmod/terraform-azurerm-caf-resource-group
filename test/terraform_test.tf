provider "azurerm"{
  version = "1.38.0"
}

module "resource_group" {
  source = "../"
  prefix = "unittest"
  resource_groups = {
    apim          = {
                    name     = "-apim-demo"
                    location = "southeastasia"
                    },
    networking    = {
                    name     = "-networking-demo"
                    location = "southeastasia"
                    },
    insights      = {
                    name     = "-insights-demo"
                    location = "southeastasia"
                    },
  }

  tags = {
    purpose = "unitTest"
    sample  = "sampleTag"
  }
}

