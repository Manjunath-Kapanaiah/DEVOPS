variable "project_id" {
  description = "GCP project ID (passed by Jenkins as TF_VAR_project_id)."
  type        = string
}

variable "region" {
  description = "GCP region."
  type        = string
  default     = "us-central1"
}
