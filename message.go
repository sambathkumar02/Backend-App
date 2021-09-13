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
