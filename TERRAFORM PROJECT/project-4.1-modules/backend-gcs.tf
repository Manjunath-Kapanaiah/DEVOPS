# ELEVATE (Task): store Terraform state in a REMOTE backend instead of a local
# terraform.tfstate file. create the bucket once, then run `terraform init` to migrate state.

terraform {
  backend "gcs" {
    bucket = "likith-tfstate-capstone" # auto assign sub paths for workspaces
    prefix = "terraform/state"
  }
}
