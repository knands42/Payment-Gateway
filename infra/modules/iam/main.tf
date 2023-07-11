resource "google_project_iam_member" "gke_cluster_creator" {
  project = var.project_id
  role    = "roles/container.admin"
  member  = "user:${var.user}"
}