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

resource "aws_cloudwatch_log_group" "tfrefresh" {
  name              = "/aws/lambda/tfrefresh"
  retention_in_days = 7
}
