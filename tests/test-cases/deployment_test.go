package parser

import (
	"testing"

	"github.com/prabdeb/go-utils/openstack/app/parser"
)

func TestDeploymentLoad(t *testing.T) {
	parser := parser.New()
	err := parser.Deployment("../test-data/deployment.yaml")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetImage(t *testing.T) {
	parser := parser.New()
	err := parser.Deployment("../test-data/deployment.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	if parser.GetImage() != "prabdeb/simple-go-server:latest" {
		t.Error("did not get expected image")
	}
}

func TestGetImageForContainer(t *testing.T) {
	parser := parser.New()
	err := parser.Deployment("../test-data/deployment.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	if parser.GetImageForContainer("simple-go-server1") != "prabdeb/simple-go-server1:latest" {
		t.Error("did not get expected image")
	}
}

func TestSetImage(t *testing.T) {
	parser := parser.New()
	err := parser.Deployment("../test-data/deployment.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	parser.SetImage("prabdeb/simple-go-server:test")
	if parser.GetImage() != "prabdeb/simple-go-server:test" {
		t.Error("did not set expected image")
	}
}

func TestSetImageForContainer(t *testing.T) {
	parser := parser.New()
	err := parser.Deployment("../test-data/deployment.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	parser.SetImageForContainer("prabdeb/simple-go-server1:test", "simple-go-server1")
	if parser.GetImageForContainer("simple-go-server1") != "prabdeb/simple-go-server1:test" {
		t.Error("did not set expected image")
	}
}
