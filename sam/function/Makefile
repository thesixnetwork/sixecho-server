.PHONY: deps clean build
	
build:
	cd contract_client && npm install && docker run --rm -v "${PWD}":/var/task lambci/lambda:build-nodejs10.x npm rebuild

deploy:
	sam package --template-file template.yaml --s3-bucket ${BUCKET_NAME} --output-template-file packaged.yaml --profile ${AWS_PROFILE}
	aws cloudformation deploy --template-file ./packaged.yaml --stack-nam wallet --capabilities CAPABILITY_IAM --profile ${AWS_PROFILE}

remove:
	aws cloudformation delete-stack --stack-name SixEchoFunction --profile ${AWS_PROFILE}
