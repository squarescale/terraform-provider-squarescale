provider "squarescale" {
  endpoint        = "https://www.nledez-dev.sqsc.io"
}

resource "squarescale_project" "demo_project" {
  name = "demo-terraform-1"
}

resource "squarescale_db" "db" {
  project = "${squarescale_project.demo_project.name}"
  engine  = "postgres"
  size    = "dev"
}

resource "squarescale_env" "demo_environnement" {
  project = "${squarescale_project.demo_project.name}"
  environnement {
    "RABBITMQ_HOST" = "rabbitmq.service.consul",
    "NODE_ENV" = "production",
  }
}

resource "squarescale_image" "rabbimq_image" {
  project   = "${squarescale_project.demo_project.name}"
  name      = "rabbitmq"
  instances = 1
}

resource "squarescale_image" "app_image" {
  project   = "${squarescale_project.demo_project.name}"
  name      = "squarescale/sqsc-demo-app"
  instances = 1
}

resource "squarescale_image" "worker_image" {
  project   = "${squarescale_project.demo_project.name}"
  name      = "squarescale/sqsc-demo-worker"
  instances = 2
}

resource "squarescale_lb" "lb" {
  project   = "${squarescale_project.demo_project.name}"
  container = "${squarescale_image.app_image.name}"
  port      = 3000
}
