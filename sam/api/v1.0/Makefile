build:
	echo "[install]" > ~/.pydistutils.cfg
	echo "prefix=" >> ~/.pydistutils.cfg
	pip2 install datasketch==1.4.3 -t ./digest_checker
	pip2 install mysql-connector-python==8.0.16  -t ./digest_checker
	pip2 install pycountry==18.12.8  -t ./digest_checker
	pip2 install SQLAlchemy==1.3.3  -t ./digest_checker
	pip2 install six==1.11.0  -t ./digest_checker
	cd ./digest_checker && rm -Rf numpy* && cd ..
	curl -o ./digest_checker/numpy-1.16.4-cp27-cp27mu-manylinux1_x86_64.whl https://files.pythonhosted.org/packages/1f/c7/198496417c9c2f6226616cff7dedf2115a4f4d0276613bab842ec8ac1e23/numpy-1.16.4-cp27-cp27mu-manylinux1_x86_64.whl 
	pip2 install ./digest_checker/numpy-1.16.4-cp27-cp27mu-manylinux1_x86_64.whl -t ./digest_checker 
	rm ./digest_checker/numpy-1.16.4-cp27-cp27mu-manylinux1_x86_64.whl
deploy:
	aws s3 cp swagger.yml s3://${BUCKET_NAME}/swagger.yml --profile ${AWS_PROFILE}
	sam package --template-file tmp-template.yaml --s3-bucket ${BUCKET_NAME} --output-template-file packaged.yaml --profile ${AWS_PROFILE}
	aws cloudformation deploy --template-file ./packaged.yaml --stack-name sixechoAPIv100 --capabilities CAPABILITY_IAM --profile ${AWS_PROFILE}
remove:
	aws cloudformation delete-stack --stack-name sixechoAPIv100 --profile ${AWS_PROFILE}
update:
	cd digest_checker && zip -r ../myDeploymentPackage.zip .
	aws lambda update-function-code --function-name sixechoAPIv100-DigestCheckerFunction-1LW2QBL2TF60Q --zip-file fileb://myDeploymentPackage.zip
