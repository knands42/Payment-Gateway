module "iam" {
  source     = "./modules/iam"
  project_id = var.project_id
  user       = var.user
}

module "gke" {
  source       = "./modules/gke"
  cluster_name = var.cluster_name
  region       = var.region
  zone         = var.zone
  pool_name    = var.pool_name
}