package main

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os/exec"
)

//Function to get docker version
func dockerVersion() (string, error) {

	cmd := exec.Command("docker", "--version")
	var command_output bytes.Buffer
	cmd.Stdout = &command_output
	err := cmd.Run()
	if err != nil {
		return "", errors.New("unable to get the docker verison")
	}

	return command_output.String(), nil
}

//Function for Deploy and Image
func deployImage(image string) (string, error) {

	//Requires sudo
	cmd := exec.Command("sudo", "docker", "run", "-d", image)
	var command_output bytes.Buffer
	cmd.Stdout = &command_output

	err := cmd.Run()
	log.Println(command_output.String())
	if err != nil {
		return "", errors.New("unable to Deploy the Image")
	}

	return command_output.String(), nil

}

type Server struct {
	UnimplementedDockerServiceServer //Denote the type corresponds to grpc DockerService service
}

//Method corresponds to the procedure call -> rpc GetDockerVersion
func (s *Server) GetDockerVersion(ctx context.Context, client *Client) (*DockerVersion, error) {

	log.Print("Request from:", client.Client)

	//Get docker version
	docker_version, err := dockerVersion()
	if err != nil {
		log.Print("Unable to find docker version")
		return nil, errors.New("obtaining docker version failed")
	}

	result := DockerVersion{
		Version: docker_version,
	}

	return &result, nil

}

//Method Denotes DeployImage method of docker service
func (s *Server) DeployImage(ctx context.Context, deploymentParams *DeploymentParams) (*Conatiner, error) {

	log.Print("Request from:", deploymentParams.Client)
	ContainerId, err := deployImage(deploymentParams.Image)
	if err != nil {
		log.Println("Unable to deploy container!")
		return nil, errors.New("deploy image failed")
	}
	result := Conatiner{
		ContainerID: ContainerId,
	}

	return &result, nil

}
