provider "squarescale" {
  endpoint        = "https://www.nledez-dev.sqsc.io"
}

resource "squarescale_project" "demo_ha_1" {
  name = "demo-ha-1"
}

resource "squarescale_db" "demo_ha_1_db" {
  project = "${squarescale_project.demo_ha_1.name}"
  engine  = "postgres"
  size    = "dev"
}

resource "squarescale_env" "demo_ha_1_rabbitmq" {
  project = "${squarescale_project.demo_ha_1.name}"
  key     = "RABBITMQ_HOST"
  value   = "rabbitmq.service.consul"
}

resource "squarescale_env" "demo_ha_1_node_env" {
  project = "${squarescale_project.demo_ha_1.name}"
  key     = "NODE_ENV"
  value   = "production"
}

resource "squarescale_image" "demo_ha_1_rabbimq_image" {
  project   = "${squarescale_project.demo_ha_1.name}"
  name      = "rabbitmq"
  instances = 1
}

resource "squarescale_image" "demo_ha_1_app_image" {
  project   = "${squarescale_project.demo_ha_1.name}"
  name      = "squarescale/sqsc-demo-app"
  instances = 1
}

resource "squarescale_image" "demo_ha_1_worker_image" {
  project   = "${squarescale_project.demo_ha_1.name}"
  name      = "squarescale/sqsc-demo-worker"
  instances = 2
}

resource "squarescale_lb" "demo_ha_1_lb" {
  project   = "${squarescale_project.demo_ha_1.name}"
  container = "squarescale/sqsc-demo-app"
  port      = 3000
}
