locals {
    prefix = "test"
    resource_groups = {
        apim = { 
            name     = "-apim-demo"
            location = "southeastasia" 
        },
        networking = {    
            name     = "-networking-demo"
            location = "southeastasia" 
        },
        insights = { 
            name     = "-insights-demo"
            location = "southeastasia" 
            tags = {
                environment     = "QA"
                owner           = "Arnaud"
                deployment      = "TerraformTest"
            }
        },
    }
    tags = {
        environment     = "DEV"
        owner           = "CAF"
    }
}