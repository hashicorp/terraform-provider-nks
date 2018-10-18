resource "nks_solution" "prometheus" {
  org_id     = "${data.nks_keysets.keyset_default.org_id}"
  cluster_id = "${nks_cluster.terraform-cluster.id}"
  solution   = "prometheus"
}
