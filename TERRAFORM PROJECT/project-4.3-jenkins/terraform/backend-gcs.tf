# Remote state in GCS for the Jenkins-managed infrastructure.
# Create the bucket once, then run `terraform init -migrate-state`.

terraform {
  backend "gcs" {
    bucket = "likith-tfstate-capstone"
    prefix = "terraform/jenkins"
  }
}
