# Remote state in GCS. Each workspace (dev/staging/prod) gets its own sub-path
# under this prefix automatically, so state stays isolated per environment.
# Create the bucket once, then run `terraform init -migrate-state`.

terraform {
  backend "gcs" {
    bucket = "likith-tfstate-capstone"
    prefix = "terraform/workspaces" # workspaces get sub-paths automatically
  }
}
