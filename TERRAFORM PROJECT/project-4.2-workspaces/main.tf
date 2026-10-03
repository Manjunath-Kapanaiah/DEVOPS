# Project 4.2 — ONE configuration, MANY environments via workspaces.
# terraform.workspace holds the current workspace name (dev/staging/prod).
# We use it to
# (a) name resources per-environment 
# (b) pick per-env settings, while each workspace keeps its own isolated state file automatically.

locals {
  env          = terraform.workspace
  allowed_envs = ["dev", "staging", "prod"]

  env_settings = {
    default = { versioning = false }
    dev     = { versioning = false }
    staging = { versioning = true }
    prod    = { versioning = true }
  }

  cfg = local.env_settings[local.env]
}

# Fails the run immediately if someone selects an unexpected workspace.
resource "null_resource" "validate_workspace" {
  lifecycle {
    precondition {
      condition = contains(local.allowed_envs, local.env)
      error_message = "Workspace '${local.env}' is not allowed. Use one of: ${join(", ",
      local.allowed_envs)}."
    }
  }
}


module "storage" {
  source     = "../modules/gcs-bucket"
  name       = "${var.project_id}-ws-${local.env}-data"
  location   = var.bucket_location
  versioning = local.cfg.versioning
  labels = {
    environment = local.env
    managed_by  = "terraform"
    project     = "4-2-workspaces"
  }
}
