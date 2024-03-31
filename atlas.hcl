variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

data "template_dir" "migrations" {
  path = "db/migrations"
  vars = {}
}

env "turso" {
  url     = "libsql+ws://127.0.0.1:8080?authToken=${var.token}"
  exclude = ["_litestream*"]
}
