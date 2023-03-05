package main

import "os"

func LoadConfig() {
	// Project ID on the project page. Example: 44066311
	os.Setenv("PROJECT_ID", "44066311")
	/*
		Open the file path in Job and paste everything after /artifacts/file/
		Example URL
		https://gitlab.com/conradgg/testCI/-/jobs/3876871442/artifacts/file/build/libs/testCI.jar
		Paste build/libs/testCI.jar
	*/
	os.Setenv("FOLDER_PATH", "build/libs/testCI.jar")
	/*
		Create a new token in the Project Folder -> Settings -> Access Tokens
		Create a new token with API scope and Maintainer role
	*/
	os.Setenv("PROJECT_TOKEN", "glpat-kopMPEg_-cTUmCtzSbbu")
	//
	os.Setenv("ENDPOINT", "/home/conradgg/file.jar")
	/*
		Create a new http webhook with the URL
		http://YourIP(Domain):port/listenpath
		Add a Job events trigger and disable SSL Verification
	*/
	os.Setenv("LISTEN_PORT", "8080")
	os.Setenv("LISTEN_PATH", "webhook")
}
