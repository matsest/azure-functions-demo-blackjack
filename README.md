# Azure Functions Demo

[![Deploy](https://github.com/matsest/azure-functions-demo/actions/workflows/deploy-function-app.yml/badge.svg)](https://github.com/matsest/azure-functions-demo/actions/workflows/deploy-function-app.yml)

This repo contains code to deploy an Azure Function in a declarative fashion. The following tools are used:

- :gear: [GitHub Actions Workflow](https://docs.github.com/en/actions/quickstart): End-to-end deployment and build of all components
- :muscle: [Bicep](https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep): infrastructure-as-code for Azure Resources
- :zap: [Azure Functions](https://docs.microsoft.com/en-us/azure/azure-functions): source code for the function in golang

The deployed function is a simple demo function that plays a game of Black Jack between two players when called and prints out the result.

```bash
$ curl https://<functionapp_name>.azurewebsites.net/api/blackjack

# Example output
cards: [S3 CJ DA D6 D7 D10 S2 CQ S7 D5 C6 SQ DK C5 H9 H2 H5 HK S5 DQ HQ C2 CA HJ HA C9 S9 D9 D2 D8 S4 CK H10 SA H8 D3 C8 SK C10 H3 DJ S6 C4 H7 H6 C3 H4 C7 S10 S8 D4 SJ]
points: 380
Mats: [S3 CJ D7] Sum: 20
Magnus: [DA D6 D10] Sum: 27
Winner: Mats
```

## Requirements

- An Azure subscription with a resource group deployed
- A service principal with Contributor permissions to the resource group
  - Add credentials as repository secrets to `AZURE_CREDENTIALS`

<details>

```bash
# Set up az cli: https://docs.microsoft.com/en-us/cli/azure/get-started-with-azure-cli

$ az group create -l {location} -n {resource group}

$ az ad sp create-for-rbac --name "myApp" --role contributor \
                         --scopes /subscriptions/{subscription-id}/resourceGroups/{resource-group}
                         --sdk-auth

# Replace {subscription-id}, {resource-group}, and {app-name} with
# the names of your subscription, resource group, and Azure function app.

# The command should output a JSON object similar to this:

{
  "clientId": "<GUID>",
  "clientSecret": "<GUID>",
  "subscriptionId": "<GUID>",
  "tenantId": "<GUID>",
  (...)
}

# Copy this and add as a repository secret named AZURE_CREDENTIALS
```

</details>

## Development

### Azure Resources

Deployment of Azure resources for Azure Functions is based on the guide found [here](https://docs.microsoft.com/en-us/azure/azure-functions/functions-infrastructure-as-code) - converted to Bicep.

To develop with Bicep install the [necessary tooling](https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep/install) (VS Code Extension, Bicep CLI).

Remember to edit the [parameter file](./bicep/main.parameters.json) before deploying.

### Functions

To develop the function, open the [function directory](./function-go) as a workspace in VS Code. This allows to use the Functions extensions to work on a local project.

This is based on the guide found [here](https://docs.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Clinux). Remember to install the necessary tooling.

#### Run function locally

Build the function:

```bash
cd function-go/src
go build -o main .
```

Run the function:

```bash
## Test the package and packaged function using Azure Functions Core Tools
cd function-go
func start                        # in one terminal
curl localhost:7071/api/blackjack # in another terminal

## Alternatively run the binary without using Azure Functions Core Tools (to quickly test the go package)
cd function-go
./main                            # in one terminal
curl localhost:8080/api/blackjack # in another terminal
```

Follow the [guide](https://docs.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Clinux) for a more descriptive guide.

#### Run the deployed function

```bash
curl https://<functionapp_name>.azurewebsites.net/api/blackjack
```
