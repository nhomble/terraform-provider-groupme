terraform {
  required_providers {
    groupme = {
      # this provider isn't in a registry yet, so I have to copy build into local terraform.d/
      source = "local.build/nhomble/groupme"
      version = "1.0.0"
    }
  }
}