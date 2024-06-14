terraform {
  required_version = "1.5.3"

  required_providers {
    confluentacl = {
      version = "0.1.1"
      source  = "brezam/confluentacl"
    }
  }
}

provider "confluentacl" {}
