variable "name" {
  description = "Name of the GCS bucket."
  type        = string
}

variable "location" {
  description = "Bucket location (e.g. US, EU, or a region like US-CENTRAL1)."
  type        = string
  default     = "US"
}

variable "storage_class" {
  description = "Storage class for the bucket."
  type        = string
  default     = "STANDARD"
}

variable "versioning" {
  description = "Enable object versioning."
  type        = bool
  default     = false
}

variable "force_destroy" {
  description = "Allow Terraform to delete the bucket even if it still has objects."
  type        = bool
  default     = true
}

variable "labels" {
  description = "Labels applied to the bucket."
  type        = map(string)
  default     = {}
}
variable "lifecycle_age_days" {
  description = "Number of days before objects are deleted"
  type        = number
  default     = 0
}

variable "kms_key_name" {
  description = "Optional Cloud KMS key used to encrypt the bucket"
  type        = string
  default     = ""
}
