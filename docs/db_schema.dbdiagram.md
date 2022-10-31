```dbdiagram
project go_auth {
  database_type: "PostgreSQL"
}

table user as U {
  id string [PK]
  username varchar(50) [not null, unique]
  email varchar(50)  [not null, unique]
  password varchar(255) [not null, unique]
  fullname varchar(255)
  created_at integer [not null]
  updated_at integer [not null]
}
```
