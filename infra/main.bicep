targetScope = 'subscription'

@description('Name of the environment (e.g., dev, prod)')
param environmentName string

@description('Azure region (e.g., eastus)')
param location string

@description('Tags to apply to all resources')
param tags object = {}

var resourceSuffix = take(uniqueString(subscription().id, environmentName, location), 6)
var resourceGroupName = 'rg-api-${environmentName}-${resourceSuffix}'

resource resourceGroup 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: resourceGroupName
  location: location
  tags: union(tags, {
    'azd-env-name': environmentName
  })
}

module containerRegistry './modules/container-registry.bicep' = {
  name: 'acr'
  scope: resourceGroup
  params: {
    name: 'acr-api'
    location: location
    tags: union(tags, {
      'azd-service-name': 'acr'
    })
  }
}

module keyVault './modules/keyvault.bicep' = {
  name: 'keyvault'
  scope: resourceGroup
  params: {
    name: 'kv-api'
    location: location
    tags: union(tags, {
      'azd-service-name': 'keyvault'
    })
  }
}

module containerApp './modules/container-app.bicep' = {
  name: 'containerapp'
  scope: resourceGroup
  params: {
    name: 'api'
    location: location
    tags: union(tags, {
      'azd-service-name': 'api'
    })
    keyVaultId: keyVault.outputs.keyVaultId
    acrLoginServer: containerRegistry.outputs.acrLoginServer
    imageName: 'portfolio-instruments-api'
    imageTag: 'latest'
  }
}

module postgres './modules/postgres.bicep' = {
  name: 'postgres'
  scope: resourceGroup
  params: {
    name: 'db'
    location: location
    tags: union(tags, {
      'azd-service-name': 'postgres'
    })
    keyVaultId: keyVault.outputs.keyVaultId
  }
}

output acrLoginServer string = containerRegistry.outputs.acrLoginServer
output acrName string = containerRegistry.outputs.acrName