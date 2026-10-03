variable "project_id" {
  description = "Your GCP project ID."
  type        = string
}

variable "region" {
  description = "Default GCP region."
  type        = string
  default     = "us-central1"
}

variable "bucket_location" {
  description = "Location for the GCS bucket."
  type        = string
  default     = "US"
}
