## Requirements
The only requirement for this lab is *Docker*. The lab will be done in a container that will have all of the tools you need.

If you wish to complete this lab without using Docker you will need to install the following:
- [Go 1.10](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/introduction.html)
- [Terraform](https://www.terraform.io/intro/getting-started/install.html)
- [Node](https://nodejs.org/en/)

## Getting Setup
The following steps will take you through how to get setup to start the lab.

### Step One: Clone this repository and change in the folder that was created:
`git clone https://github.com/gowlg/lab-custom-terraform-provider.git && cd lab-custom-terraform-provider`

### Step Two: Build the Docker image for the lab
`docker build -t terraform-provider-lab .`

### Step Three: Run the container mounting the repository directory and mapping port 3000 to port 3000 of the container
`docker run -it -v $PWD:/go/src/github.com/gowlg/lab-custom-terraform-provider -w /go/src/github.com/gowlg/lab-custom-terraform-provider -p 3000:3000 terraform-provider-lab bash`

You will now have a bash session inside the lab's docker container.

### Step Five: Start the API
```
cd ./api
npm install
npm run start &
cd ..
```

### Step Five: Open your favorite editor and get hacking
All go and terraform commands will be run in the the container.
You can edit the files in an editor outside the container and it will effect the same files inside the container.


`go build -o ./test/terraform-provider-uws`
`cd test`
`terraform init`
`terraform apply`

