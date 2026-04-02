echo ">>> Attaching ACR to AKS..."
az aks update \
  --name $AKS_NAME \
  --resource-group $RESOURCE_GROUP \
  --attach-acr cloudopsacrdev
