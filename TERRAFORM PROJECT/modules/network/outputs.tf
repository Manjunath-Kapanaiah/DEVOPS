output "network_name" {
  description = "The VPC network name."
  value       = google_compute_network.this.name
}

output "network_self_link" {
  description = "The VPC network self link."
  value       = google_compute_network.this.self_link
}

output "subnet_name" {
  description = "The subnet name."
  value       = google_compute_subnetwork.this.name
}
