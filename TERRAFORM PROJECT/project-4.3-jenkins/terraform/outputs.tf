output "bucket_name" {
  description = "Bucket created by the Jenkins pipeline."
  value       = google_storage_bucket.pipeline.name
}
