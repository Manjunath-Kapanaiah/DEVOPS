variable "project_id" {
  description = "capstone-project-505907"
  type        = string
}

variable "region" {
  description = "Default Region"
  type        = string
  default     = "us-central1"
}

variable "environment" {
  description = "Environment label used in resource names."
  type        = string
  default     = "dev"
}

variable "bucket_location" {
  description = "Location for the GCS bucket (US, EU, or a region)."
  type        = string
  default     = "US"
}
