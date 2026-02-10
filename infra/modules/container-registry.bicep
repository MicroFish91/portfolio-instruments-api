targetScope = 'resourceGroup'

@description('Registry name for Azure Container Registry')
param name string
@description('Location for registry')
param location string = resourceGroup().location
@description('Tags for registry')
param tags object = {}

var resourceSuffix = take(uniqueString(subscription().id, resourceGroup().name, name), 6)
var registryName = '${name}-${resourceSuffix}'

resource acr 'Microsoft.ContainerRegistry/registries@2023-01-01-preview' = {
  name: registryName
  location: location
  tags: tags
  sku: {
    name: 'Basic'
  }
  properties: {
    adminUserEnabled: false
    publicNetworkAccess: 'Enabled'
  }
}

output acrId string = acr.id
output acrName string = acr.name
output acrLoginServer string = acr.properties.loginServer