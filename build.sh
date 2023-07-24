export CODE_IMAGE_REPO_NAME=go-api
export VERSION=$(cat semver)
export AWS_ACCOUNT_ID='409783107836' #infrastructure account
echo building $VERSION of $CODE_IMAGE_REPO_NAME
echo Build started on `date`

echo Logging in to Amazon ECR...
aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com

echo checking for lambda docker repo...
exists=$(aws ecr describe-repositories --output json | jq --arg CODE_IMAGE_REPO_NAME $CODE_IMAGE_REPO_NAME '.repositories  | map(.repositoryName==$CODE_IMAGE_REPO_NAME) | any')

if [[ $exists == "false" ]]; 
    then aws ecr create-repository --repository-name $CODE_IMAGE_REPO_NAME; 
fi; 

echo Building the lambda docker image...          
docker build -t $CODE_IMAGE_REPO_NAME:$VERSION ./src
docker tag $CODE_IMAGE_REPO_NAME:$VERSION $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$CODE_IMAGE_REPO_NAME:$VERSION

echo Pushing the lambda docker image...
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$CODE_IMAGE_REPO_NAME:$VERSION