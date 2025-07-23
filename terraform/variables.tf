variable "cloud_id" {}
variable "folder_id" {}
variable "oauth_token" {}
variable "zone" { default = "ru-central1-a" }
variable "image_id" { default = "fd8u1vhsbjt0g2g4fjit" } # Ubuntu 22.04 LTS, можно обновить при необходимости
variable "subnet_id" {}
variable "ssh_public_key_path" {}
variable "ssh_public_key" {}