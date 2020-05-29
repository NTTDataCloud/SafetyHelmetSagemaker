# docprac
Docker practice in golang

## [Docker Commends](#docker)
## [Deploy to Fargate](#deploy-to-fargate)

## Docker

### Image deletion 

Remove all containers along with its storage. 

```
% docker rm -vf $(docker ps -a -q)
```

Remove all local images.

```
% docker rmi -f $(docker images -a -q)
```

### Image Build 
```
% docker build -t <name of image>:<version> .
```
### Create ECR Repository
```
% aws ecr create-repository \
    --repository-name <repo name> \
    --image-scanning-configuration scanOnPush=true \
    --region <region>
```

### List Images
```
% docker images
```

### Run Image
```
% docker run -t -i -p <output port>:<input port> <image name>
```

### Image Tag
```
% docker tag <local docker image name>:<local docker image version> <aws account id>:dkr.ecr.<region>.amazonaws.com/<ecr docker image name>:<ecr docker version>
```

### Image Push to ECR
```
% docker push <aws account id>:dkr.ecr.<region>.amazonaws.com/<ecr docker image name>:<ecr docker version>
```

### Image Pull ECR Image Down
```
% docker pull <aws account id>:dkr.ecr.<region>.amazonaws.com/<ecr docker image name>:<ecr docker version>
```

### Delete ECR Image
```
% aws ecr batch-delete-image --repository-name <ecr docker image name> --image-ids imageTag=<ecr docker version>
```

### Docker Push Authentication 
```
% aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com 
```
---
## Deploy To ECR

### Step 1. clone repo
```
% git clone git@github.com:avachen2005/docprac.git
% cd docprac
```
### Step 2. [Authenticate docker with IAM](#docker-push-authentication) 
### Step 3. [Build docker image](#image-build)
### Step 4. [Tag docker image](#image-tag)
### Step 5. [Push Image to ECR](#image-push-to-ecr)

---
## Renew Task
## Renew Service