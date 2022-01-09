resource "aws_cloudwatch_event_rule" "event" {
  name = "event"
  event_pattern = jsonencode({
    account = [data.aws_caller_identity.current.account_id]
  })
}

data "aws_lambda_function" "tfrefresh" {
  function_name = "tfrefresh"
}

resource "aws_cloudwatch_event_target" "tfrefresh" {
  arn  = data.aws_lambda_function.tfrefresh.arn
  rule = aws_cloudwatch_event_rule.event.name
}

resource "aws_lambda_permission" "tfrefresh" {
  function_name = data.aws_lambda_function.tfrefresh.function_name
  action        = "lambda:InvokeFunction"
  source_arn    = aws_cloudwatch_event_rule.event.arn
  principal     = "events.amazonaws.com"
}
