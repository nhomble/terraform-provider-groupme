terraform {
  required_providers {
    groupme = {
      # this provider isn't in a registry yet, so I have to copy build into local terraform.d/
      source = "local/providers/groupme"
      version = "1.0.0"
    }

    random = {
      source = "registry.terraform.io/hashicorp/random"
      version = "3.1.0"
    }
  }
}
