schema "main" {
}

table "users" {
  schema = schema.main
  column "id" {
    null = false
    type = integer
    auto_increment = true
  }
  primary_key  {
    columns = [column.id]
  }
  index "idx_id" {
    columns = [ 
      column.id
     ]
  }

  column "username" {
    type = text
    null = false
  }

  column "email" {
    type = text
    null = false
  }

  column "salted_hash" {
    type = text
    null = false
  }

  column "firstname" {
    type = text
    null = false
  }

  column "lastname" {
    type = text
    null = false
  }
}
