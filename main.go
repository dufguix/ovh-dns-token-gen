package main

import (
	"fmt"
	"strconv"

	"github.com/ovh/go-ovh/ovh"
)

var client *ovh.Client

func main() {
	// Create a client using credentials from config files or environment variables
	var err error
	client, err = ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	for {
		if printMenu() {
			break
		}
	}

}

// return true to exit
func printMenu() bool {
	fmt.Println("\n\nMenu:")
	fmt.Println("1) See all applications infos")
	fmt.Println("2) See all credentials infos")
	fmt.Println("3) See credentials of an app")
	fmt.Println("4) Create credentials for DNS zone API")
	fmt.Println("5) Delete an app and all its credentials")
	fmt.Println("6) Delete a credential")
	fmt.Println("7) Exit")
	fmt.Print("Enter a menu number: ")
	var menuAnswer string
	_, err := fmt.Scanln(&menuAnswer)
	if err != nil {
		fmt.Println(err)
		return true
	}
	switch menuAnswer {
	case "1":
		apps := getApplications()
		for _, app := range apps {
			printApplicationInfo(app)
		}
	case "2":
		creds := getAllCredentials()
		for _, cred := range creds {
			printCredentialsInfo(cred)
		}
	case "3":
		fmt.Print("Enter the app ID: ")
		var idAnswer string
		_, err := fmt.Scanln(&idAnswer)
		if err != nil {
			fmt.Println(err)
			return true
		}
		id, err := strconv.Atoi(idAnswer)
		if err != nil {
			fmt.Println("This is not a number. Try again.")
			break
		}
		creds := getCredentials(id)
		if len(creds) == 0 {
			fmt.Println("No credential.")
		}
		for _, cred := range creds {
			printCredentialsInfo(cred)
		}
	case "4":
		fmt.Print("Enter your involved domain (ex: example.com): ")
		var domainAnswer string
		_, err := fmt.Scanln(&domainAnswer)
		if err != nil {
			fmt.Println(err)
			return true
		}
		generateConsumerKey(domainAnswer)
	case "5":
		fmt.Print("Enter the application ID: ")
		var idAnswer string
		_, err := fmt.Scanln(&idAnswer)
		if err != nil {
			fmt.Println(err)
			return true
		}
		id, err := strconv.Atoi(idAnswer)
		if err != nil {
			fmt.Println("This is not a number. Try again.")
			break
		}
		deleteApplication(id)
	case "6":
		fmt.Print("Enter the credential ID: ")
		var idAnswer string
		_, err := fmt.Scanln(&idAnswer)
		if err != nil {
			fmt.Println(err)
			return true
		}
		id, err := strconv.Atoi(idAnswer)
		if err != nil {
			fmt.Println("This is not a number. Try again.")
			break
		}
		deleteCredential(id)
	case "7":
		fmt.Println("Bye")
		return true
	default:
		println("Invalid menu. Try again.")
	}
	return false
}

func getApplications() []application {
	applicationIds := []int{}
	err := client.Get("/me/api/application", &applicationIds)
	if err != nil {
		fmt.Println(err)
		return []application{}
	}

	applications := make([]application, 0, len(applicationIds))
	for _, appId := range applicationIds {
		app := application{}
		err := client.Get("/me/api/application/"+strconv.Itoa(appId), &app)
		if err != nil {
			fmt.Println(err)
			continue
		}
		applications = append(applications, app)
	}
	return applications
}

func printApplicationInfo(app application) {
	fmt.Printf("App ID: %v\n", app.ID)
	fmt.Printf("\tname: %v\n", app.Name)
	fmt.Printf("\tdescription: %v\n", app.Description)
	fmt.Printf("\tstatus: %v\n", app.Status)
	fmt.Printf("\tkey: %v\n", app.Key)
}

func getAllCredentials() []credential {
	credentialIds := []int{}
	err := client.Get("/me/api/credential/", &credentialIds)
	if err != nil {
		fmt.Println(err)
		return []credential{}
	}

	credentials := make([]credential, 0, len(credentialIds))
	for _, credId := range credentialIds {
		cred := credential{}
		err := client.Get("/me/api/credential/"+strconv.Itoa(credId), &cred)
		if err != nil {
			fmt.Println(err)
			continue
		}
		credentials = append(credentials, cred)
	}
	return credentials
}

func getCredentials(appId int) []credential {
	credentialIds := []int{}
	err := client.Get("/me/api/credential?applicationId="+strconv.Itoa(appId), &credentialIds)
	if err != nil {
		fmt.Println(err)
		return []credential{}
	}

	credentials := make([]credential, 0, len(credentialIds))
	for _, credId := range credentialIds {
		cred := credential{}
		err := client.Get("/me/api/credential/"+strconv.Itoa(credId), &cred)
		if err != nil {
			fmt.Println(err)
			continue
		}
		credentials = append(credentials, cred)
	}
	return credentials
}

func printCredentialsInfo(cred credential) {
	fmt.Printf("Cred ID: %v\n", cred.ID)
	fmt.Printf("\tAppID: %v\n", cred.AppID)
	fmt.Printf("\tStatus: %v\n", cred.Status)
	fmt.Printf("\tLastUse: %v\n", cred.LastUse)
	fmt.Printf("\tExpiration: %v\n", cred.Expiration)
	fmt.Printf("\tCreation: %v\n", cred.Creation)
	fmt.Printf("\tRules: %+v\n", cred.Rules)
}

// domain without root dot
func generateConsumerKey(domain string) {
	ckReq := client.NewCkRequest()
	ckReq.AddRule("POST", "/domain/zone/"+domain+"/record")
	ckReq.AddRule("POST", "/domain/zone/"+domain+"/refresh")
	ckReq.AddRule("DELETE", "/domain/zone/"+domain+"/record/*")
	ckReq.Redirection = "https://" + domain
	response, err := ckReq.Do()
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	fmt.Printf("Generated consumer key (copy it): %s\n", response.ConsumerKey)
	fmt.Printf("Please visit %s to validate it\n", response.ValidationURL)
}

func deleteApplication(id int) {
	err := client.Delete("/me/api/application/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Application %v deleted.\n", id)
}

func deleteCredential(id int) {
	err := client.Delete("/me/api/credential/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Credential %v deleted.\n", id)
}
