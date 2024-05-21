package storage

import (
	"os"
	"testing"
)

func TestNewClientFromConfig(t *testing.T) {
	openstackConfig, err := os.ReadFile(".openstack.config")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	src := string(openstackConfig)

	_, err = NewClientFromString(src)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	awsConfig, err := os.ReadFile(".aws.config")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	src = string(awsConfig)

	_, err = NewClientFromString(src)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mockConfig, err := os.ReadFile(".mock.config")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	src = string(mockConfig)

	_, err = NewClientFromString(src)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
