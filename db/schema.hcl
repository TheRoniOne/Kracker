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
    type = text(255)
    null = false
  }

  column "salted_hash" {
    type = text
    null = false
  }

  column "firstname" {
    type = text(50)
    null = false
  }

  column "lastname" {
    type = text(50)
    null = false
  }
}
