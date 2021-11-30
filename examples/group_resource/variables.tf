variable groupme_token {
  type        = string
  description = "api key to call groupme apis"
  sensitive   = true
  validation {
    condition     = length(var.groupme_token) > 8
    error_message = "Must provide a token!"
  }
}
