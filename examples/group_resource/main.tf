provider groupme {
  api_key = var.groupme_token
}

resource groupme_group all {
  name = "test-provider2"
}
