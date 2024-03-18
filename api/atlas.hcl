data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres",
  ]
}

locals {
  db_host = getenv("DB_HOST")
  db_user = getenv("DB_USER")
  db_pass = getenv("DB_PASSWORD")
  db_name = getenv("DB_NAME")
  db_port = getenv("DB_PORT")
  db_sslmode = getenv("DB_SSLMODE")
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://${local.db_user}:${local.db_pass}@${local.db_host}:${local.db_port}/${local.db_name}?sslmode=${local.db_sslmode}"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
