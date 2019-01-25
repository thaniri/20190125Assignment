# 20190125Assignment

# How this works:
0. Install awscli on your computer and have credentials set up.
1. Run this command to deploy a VPC, Autoscaling Group, S3 Bucket, and AWS CodeDeploy:
`aws cloudformation deploy --stack-name <prod/dev/qa-stack> --template-file ./vpc.yaml --region <region you want> --capabilities CAPABILITY_IAM`
2. Create IAM roles for EC2 -> S3 and CodeDeploy <-> EC2
3. Make a pull request to this repo, travis CI will run some tests on it and upload it to an S3 bucket.
4. Run this command to deploy the code: `aws deploy create-deployment --application-name HelloWorld --deployment-config-name CodeDeployDefault.OneAtATime --deployment-group-name HelloWorld --description "demo" --s3-location bucket=20190125assignment-thaniri,bundleType=tgz,key=hello_world_app.tgz --region us-east-1`
5. Hit `http://loadbalancer-39959434.us-east-1.elb.amazonaws.com:8080` to see it live.

## Directory Structure:
```
.
├── binary # Directory contains install scripts for CodeDeploy, also used by Travis CI generate a binary.
│   ├── appspec.yml # File used by CodeDeploy to install and run the binary
│   └── install_scripts
│       ├── change_permissions.sh # Makes every script executable
│       ├── fetch_hello_world.sh # Fetches the binary from S3
│       ├── start_hello_world.sh # Runs the binary as a daemon
│       ├── stop_hello_world.sh # Kills the process.
│       └── systemd_unit_file
│           └── hello_world.service # A systemd unit file to run the web app as a daemon.
├── hello_world.go # Source Code
├── README.md
├── sample_post.txt # A sample CURL command to test POST responses.
└── vpc.yaml # Cloudformation template for the whole app.
```

# Improvements that should be made given more time:
1. Refactor the CFN stack to not be one big piece of YAML. Make use of include statements.
2. Do everything IAM using cloudformation
3. Get TravisCI to trigger a CodeDeploy automatically (very easy to do): https://docs.travis-ci.com/user/deployment/codedeploy/
4. Give minimum permissions required to all resources for this to run.
5. Start and stop scripts are overly simplistic.
6. Credentials should come from Vault.
7. VPC is very simplistic. It just autoassigns a public IP to every web app to make development easy.
8. Name variables in a sane way.

## Some assumptions made:

1. I premade some IAM policies and roles to make development quicker.
2. The bucket name is the same for travis CI and S3. Which ends up meaning it was named in two places.

# Where I got stuck:
* My development practice is to build something manually and then code it afterwards. I got stuck > 1hr on an EC2 behaviour where an IAM profile attached to an EC2 instance cannot be changed after launch.

