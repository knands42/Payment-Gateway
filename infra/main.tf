module "iam" {
  source     = "./modules/iam"
  project_id = var.project_id
  user       = var.user
}

module "vpc" {
  source     = "./modules/vpc"
  project_id = var.project_id
  region     = var.region
}

module "gke" {
  source       = "./modules/gke"
  cluster_name = var.cluster_name
  region       = var.region
  zone         = var.zone
  pool_name    = var.pool_name
  subnet_name  = module.vpc.subnet_name
  vpc_name     = module.vpc.vpc_name
}