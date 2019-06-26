package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"os"
	"context"
//	"log"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	
)

const fileSecret = "/kvmnt"
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Azure AD Pod Identity + Keyvault Flex Volume mount demo")
}

func getFileContent(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(fileSecret)
	if err != nil {
		io.WriteString(w, "Provided file is not mounted, check your Identity")
	} else {
		io.WriteString(w, string(data))
	}
	
}

func getKeyvaultSecret(w http.ResponseWriter, r *http.Request) {
	keyvaultName := os.Getenv("AZURE_KEYVAULT_NAME")
	keyvaultSecretName := os.Getenv("AZURE_KEYVAULT_SECRET_NAME")

	keyClient := keyvault.New()
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	if err != nil {
		io.WriteString(w,"failed to retrieve the authorization")
		return
	} else {
		keyClient.Authorizer = authorizer
	}

	secret, err := keyClient.GetSecret(context.Background(), fmt.Sprintf("https://%s.vault.azure.net", keyvaultName), keyvaultSecretName, "")
	if err != nil {
		//log.Printf("failed to retrieve the Keyvault secret")
		io.WriteString(w,"failed to retrieve the Keyvault secret")
		return
	}


	io.WriteString(w, fmt.Sprintf("The value of the Keyvault secret is: %v", *secret.Value))
}

func main() {
	
	http.HandleFunc("/", hello)
	http.HandleFunc("/file", getFileContent)
	http.HandleFunc("/vaultsecret", getKeyvaultSecret)
	http.ListenAndServe(":8080", nil)
}
