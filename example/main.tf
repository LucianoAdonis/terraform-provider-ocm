provider "ocm" {
  username = "${var.username}"
  password = "${var.password}"
  domain   = "${var.domain}"
}

resource "ocm_storage" "demo" {
  name       = "${var.storageName}"
  size       = "${var.storageSize}"
  path       = "${var.storagePath}"
  properties = "${var.storageProperties}"
  bootable   = "${var.storageBootable}"

  //bootable   = "${var.storageBootable}"
  //image = "${var.storageImage}"
}