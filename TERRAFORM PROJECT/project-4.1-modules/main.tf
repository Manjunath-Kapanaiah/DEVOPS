# Project 4.1 — the ROOT (main) configuration.
# resources are not created directly, instead it CALLS reusable modules.
# Related to Tasks 2 & 3: "create a module" and "use the module in main config".

module "storage" {
  source     = "../modules/gcs-bucket"
  name       = "${var.project_id}-${var.environment}-data"
  location   = var.bucket_location
  versioning = true
  labels = {
    environment = var.environment
    owner       = "likith"
    managed_by  = "terraform"
    project     = "4-1-modules"
  }
}

module "network" {
  source = "../modules/network"
  name   = "${var.environment}-vpc"
  region = var.region
}
