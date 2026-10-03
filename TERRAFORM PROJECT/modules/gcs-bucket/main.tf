# Reusable module: a single Google Cloud Storage bucket.
# Kept generic so the root configs (4.1, 4.2, 4.3) can create many buckets from the same building block.

resource "google_storage_bucket" "this" {
  name                        = var.name
  location                    = var.location
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced" # block public exposure

  versioning { enabled = var.versioning }

  dynamic "lifecycle_rule" { # optional retention/cleanup
    for_each = var.lifecycle_age_days > 0 ? [1] : []
    content {
      condition { age = var.lifecycle_age_days }
      action { type = "Delete" }
    }
  }

  dynamic "encryption" { # optional CMEK
    for_each = var.kms_key_name != "" ? [1] : []
    content { default_kms_key_name = var.kms_key_name }
  }
}
