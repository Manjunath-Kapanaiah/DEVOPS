variable "name" {
  description = "Name of the VPC network."
  type        = string
}

variable "region" {
  description = "Region for the subnet."
  type        = string
  default     = "us-central1"
}

variable "subnet_cidr" {
  description = "CIDR range for the subnet."
  type        = string
  default     = "10.10.0.0/24"
}
