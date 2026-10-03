output "workspace" {
  description = "The active Terraform workspace (environment)."
  value       = terraform.workspace
}

output "bucket_name" {
  description = "Environment-specific bucket name."
  value       = module.storage.name
}

output "versioning_enabled" {
  description = "Whether versioning is on for this environment."
  value       = local.cfg.versioning
}
