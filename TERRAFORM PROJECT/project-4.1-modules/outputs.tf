# Task 4 — Output variables. Expose useful facts about the deployed
# infrastructure. After `terraform apply`, view them with `terraform output`.
output "bucket_name" {
  description = "Name of the created storage bucket."
  value       = module.storage.name
}

output "bucket_url" {
  description = "gs:// URL of the storage bucket."
  value       = module.storage.url
}

output "network_name" {
  description = "Name of the created VPC network."
  value       = module.network.network_name
}

output "subnet_name" {
  description = "Name of the created subnet."
  value       = module.network.subnet_name
}
