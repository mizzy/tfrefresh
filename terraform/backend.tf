terraform {
  backend "s3" {
    region         = "ap-northeast-1"
    dynamodb_table = "terraform"
    bucket         = "terraform.mizzy.org"
    key            = "tfrefresh.tfstate"
    session_name   = "tfrefresh"
    encrypt        = true
    role_arn       = "arn:aws:iam::019115212452:role/terraform"
  }
}
