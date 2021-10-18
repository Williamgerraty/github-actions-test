terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.26.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.0.1"
    }
  }
  required_version = ">= 0.14"

  backend "remote" {
    organization = "gerraty"

    workspaces {
      name = "gh-actions-demo"
    }
  }
}


provider "aws" {
  region = "ap-southeast-2"
}

resource "aws_s3_bucket" "b" {
  bucket = "willjg22727272772727"
  acl    = "private"

  versioning {
    enabled = true
  }
}