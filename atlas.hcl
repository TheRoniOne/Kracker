data "template_dir" "migrations" {
  path = "db/migrations"
  vars = {}
}

env "local" {
  url = "docker://postgres/16/dev?search_path=public"
  migration {
    dir = data.template_dir.migrations.url
  }
}