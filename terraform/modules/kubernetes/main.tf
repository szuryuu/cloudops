resource "azurerm_kubernetes_cluster" "main" {
  name                = "${var.project_name}-aks-${var.environment}"
  location            = var.resource_group_location
  resource_group_name = var.resource_group_name
  dns_prefix          = "acctestagent1"

  default_node_pool {
    name       = "default"
    node_count = 2
    vm_size    = "Standard_B1s"
  }

  identity {
    type = "SystemAssigned"
  }

  linux_profile {
    admin_username = "adminuser"

    ssh_key {
      
    }
  }

  tags = {
    Project_Name = var.project_name
    Environment  = var.environment
  }
}
