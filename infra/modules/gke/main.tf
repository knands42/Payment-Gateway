# GKE cluster
resource "google_container_cluster" "primary" {
  name     = var.cluster_name
  location = var.region

  initial_node_count = 1
  remove_default_node_pool = true
  # network = google_compute_network.vpc_network.name

  node_config {
    machine_type = "n1-standard-1"
  }
}

# GKE node pool
resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = var.pool_name
  location   = var.zone
  cluster    = google_container_cluster.primary.name
  node_count = 1
  

  node_config {
    preemptible  = true
    machine_type = "e2-micro"
    
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
}
