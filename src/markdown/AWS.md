# AWS CLI commands

## Sends Cloudwatch Event notifications to Slack using Lambda function(Golang)
Based on next articles:  [1](https://read.acloud.guru/slack-notification-with-cloudwatch-alarms-lambda-6f2cc77b463a) and [2](https://gist.github.com/rantav/c096294f6f35c45155b4)

1. Create EC2 instance

    ```bash
    aws sns create-topic --name cpu-utilization-alarm-topic --region us-east-1
    ```

2. Create a slack incoming webhook and get your token

3. Encrypt this token with KMS

    Encrypt this token with KMS and paste the CiphertextBlob at 'Replace with your KMS token' (KMS encryption is a bonus, not required, you may just use the token plaintext).

    Example how to encrypt: (replace 22e06448-f73c-42c4-b18f-74f91eb7bc1a with your own key-id from KMS - create a new key and copy the ID)
    ```bash
    $ aws kms encrypt --key-id 22e06448-f73c-42c4-b18f-74f91eb7bc1a --plaintext "your slack token"

     {
        "KeyId": "arn:aws:kms:us-west-2:xxxxxxxxxx:key/22e06448-f73c-42c4-b18f-74f91eb7bc1a",
        "CiphertextBlob": "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"
    }
    ```
4. Create Lambda function
    ```sh
    aws lambda create-function --function-name DevOpsToSlack \
    --zip-file fileb://./deployment.zip \
    --runtime go1.x --handler main \
    --role IAM_ROLE_ARN \
    --environment Variables={SLACK_WEBHOOK=https://hooks.slack.com/services/TOKEN} \
    --region us-east-1 \
    ```
