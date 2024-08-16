terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
      version = "~> 5.0"
    }
  }
}

provider "heroku" {
  api_key = var.heroku_api_key
}

resource "heroku_app" "backend" {
  name   = "todolist-project"
  region = "us"
}

resource "heroku_addon" "heroku-postgresql" {
  app_id  = heroku_app.backend.id
  plan = "heroku-postgresql:essential-0"
}