variable "username" {}
variable "password" {}
variable "domain" {}

variable "storageName" {
  default     = "vol-terraform"
}

variable "storageSize" {
  default     = "30G"
}

variable "storagePath" {
  default     = ""
}

variable "storageProperties" {
  default     = ""
}

variable "storageBootable" {
  default     = true
}

variable "storageImage" {
  default = ""
}
