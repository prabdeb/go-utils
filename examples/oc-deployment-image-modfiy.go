package examples

import (
	"log"

	"github.com/prabdeb/go-utils/openstack/app/parser"
)

func ocDeploymentImageModify() {
	parser := parser.New()
	err := parser.Deployment("app-deployment.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	parser.SetImage("prabdeb/simple-go-server:latest")
	err = parser.WriteDeployment("app-deployment.yaml.parsed")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Successfully create parsed deployment config app-deployment.yaml.parsed")
}
