# Reusable module: a VPC network + one subnet.
# A second module so Project 4.1's "multiple reusable modules" elevate item is covered, and to show modules composing real infrastructure.

resource "google_compute_network" "this" {
  name                    = var.name
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "this" {
  name          = "${var.name}-subnet"
  ip_cidr_range = var.subnet_cidr
  region        = var.region
  network       = google_compute_network.this.id
}
