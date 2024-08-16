output "app_url" {
  value = heroku_app.backend.web_url
  description = "The URL of the Heroku app"
}