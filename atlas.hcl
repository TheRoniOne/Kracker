data "template_dir" "migrations" {
  path = "db/migrations"
  vars = {}
}

locals {
  db_host = getenv("DB_HOST")
  db_port = getenv("DB_PORT")
  db_name = getenv("DB_NAME")
  db_user = getenv("DB_USER")
  db_password = urlescape(getenv("DB_PASSWORD"))
}

env "local" {
  url = "postgres://${local.db_user}:${local.db_password}@${local.db_host}:${local.db_port}/${local.db_user}?search_path=public&sslmode=disable"
  migration {
    dir = data.template_dir.migrations.url
  }
}
