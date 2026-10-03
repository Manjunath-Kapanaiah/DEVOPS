# Minimal infrastructure the Jenkins pipeline provisions. Kept self-contained
# (a single bucket, no module path) so the pipeline working directory is simple.
resource "google_storage_bucket" "pipeline" {
  name          = "${var.project_id}-jenkins-demo"
  location      = "US"
  force_destroy = true

  uniform_bucket_level_access = true

  labels = {
    managed_by = "terraform"
    pipeline   = "jenkins"
  }
}
