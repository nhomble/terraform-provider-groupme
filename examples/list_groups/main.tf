provider groupme {
  api_key = var.groupme_token
}

data groupme_groups all {

}

data groupme_group a_group {
  group_id = element(tolist(data.groupme_groups.all.ids), 0)
}

