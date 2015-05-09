package config

import (
	"os"
	"testing"
)

func TestValidateInfo(t *testing.T) {
	info := Info{
		Nick:  "Foobang",
		User:  "bar",
		Gecos: "fake info",
	}

	if !info.Validate() {
		t.Fatalf("%#v not valid? wtf", info)
	}
}

func TestValidateNetwork(t *testing.T) {
	net := Network{
		Name:         "ShadowNET",
		Host:         "127.0.0.1",
		Port:         5535,
		UseSSL:       false,
		ServicesPass: "foobang",
	}

	if !net.Validate() {
		t.Fatalf("%#v not valid? wtf", net)
	}
}

func TestLoad(t *testing.T) {
	fin, err := os.Open("./example.conf.json")
	if err != nil {
		t.Fatal(err)
	}

	defer fin.Close()

	_, err = Load(fin)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadFile(t *testing.T) {
	_, err := LoadFile("./example.conf.json")
	if err != nil {
		t.Fatal(err)
	}
}
