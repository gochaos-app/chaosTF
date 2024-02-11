# ChaosTF

This project reads terraform files and outputs go-chaos code. 

## Usage
chaosTF has a single action to read a terraform file: 
```
chaosTF read terraform/main.tf env:dev kill
```

where the file located at `terraform/main.tf` will be read and find every resource with tags `env:dev`, for the last argument, user can choose from `basic` and `kill` 

* `basic` action will only choose from stop and reboot. 
* `kill` action will choose to terminate or shutdown infrastructure.

default is `kill`

This command will create a file called output.hcl that the user can clean or modify for a  more customized approach. 