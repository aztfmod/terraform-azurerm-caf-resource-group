# rg_simple

Creates a set of test resource groups. 

## Usage
To run this example, simply execute: 

```hcl
terraform init
terraform plan
terraform apply
```

Once you are done, just run 
```hcl
terraform destroy
```

## Outputs
| Name | Description |
| --   | -- |
| object | Returns the full set of resource group objects created | 
| names | Returns a map of resource_group key -> resource_group name |
| ids | Returns a map of resource_group key -> resource_group id | 