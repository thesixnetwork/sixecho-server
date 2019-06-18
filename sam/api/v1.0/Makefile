build:
deploy:
	aws s3 cp swagger.yml s3://${BUCKET_NAME}/swagger.yml --profile ${AWS_PROFILE}
	sam package --template-file tmp-template.yaml --s3-bucket ${BUCKET_NAME} --output-template-file packaged.yaml --profile ${AWS_PROFILE}
	aws cloudformation deploy --template-file ./packaged.yaml --stack-name sixechoAPIv100 --capabilities CAPABILITY_IAM --profile ${AWS_PROFILE}
remove:
	aws cloudformation delete-stack --stack-name sixechoAPIv100 --profile ${AWS_PROFILE}
