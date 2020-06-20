To run this application we just need to install docker engine(docker) on your respective os.

Once the Docker engine is installed we just have to build the docker image and run the container with that image.
Inside the container environment go-app  would already be installed(no need to do it explicitly) so it doesn't matter whether go-app is
present on the host os or not(Advantage of using containers)

To build the image use below command from the same dir where this repo has been cloned:
sudo docker build . -t go-docker

For running the container:
sudo docker run --detach --name az go-docker:latest

For entering in the container environment
sudo docker exec -it az /bin/bash

Note:1-The string for which app has to be run is to be provided in the inputfile explicitly by the user
     2-Required output can be seen in Result.txt file (Present inside the container)


