resource "aws_iam_role" "tfrefresh" {
  name               = "tfrefresh"
  assume_role_policy = data.aws_iam_policy_document.tfrefresh_assume_role_policy.json
}

data "aws_iam_policy_document" "tfrefresh_assume_role_policy" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "tfrefresh_lambda_basic_execution_role" {
  role       = aws_iam_role.tfrefresh.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "tfrefresh_read_only" {
  role       = aws_iam_role.tfrefresh.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

data "aws_iam_policy_document" "tfrefresh_assume_terraform_role" {
  statement {
    effect    = "Allow"
    actions   = ["sts:AssumeRole"]
    resources = ["arn:aws:iam::019115212452:role/terraform"]
  }
}


resource "aws_iam_policy" "tfrefresh_assume_terraform_role" {
  policy = data.aws_iam_policy_document.tfrefresh_assume_terraform_role.json
}

resource "aws_iam_role_policy_attachment" "tfrefresh_assume_terraform_role" {
  role       = aws_iam_role.tfrefresh.name
  policy_arn = aws_iam_policy.tfrefresh_assume_terraform_role.arn
}

resource "aws_cloudwatch_log_group" "tfrefresh" {
  name              = "/aws/lambda/tfrefresh"
  retention_in_days = 7
}
