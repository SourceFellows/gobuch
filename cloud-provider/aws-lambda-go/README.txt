```
GOOS=linux go build main.go
zip function.zip main
aws lambda create-function \
    --function-name aws-lambda-sample \
    --runtime go1.x \
    --zip-file fileb://function.zip
    --handler main
    --role arn:aws:iam::<xx>:role/<xx>
```