// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedapplications_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2021-07-01/applications"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type ManagedApplicationResource struct{}

func TestAccManagedApplication_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccManagedApplication_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccManagedApplication_sequential(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"Resource": {
			"complete": testAccManagedApplication_complete,
			"update":   testAccManagedApplication_update,
		},
	}

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}

func testAccManagedApplication_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceSequentialTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func testAccManagedApplication_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceSequentialTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("ServiceCatalog"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("MarketPlace"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("kind").HasValue("ServiceCatalog"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccManagedApplication_updateParameters(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("parameters.%").HasValue("3"),
				check.That(data.ResourceName).Key("parameter_values").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.parameterValues(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("parameters.%").Exists(),
				check.That(data.ResourceName).Key("parameter_values").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("parameters.%").HasValue("3"),
				check.That(data.ResourceName).Key("parameter_values").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccManagedApplication_parameterValues(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_managed_application", "test")
	r := ManagedApplicationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.parameterValues(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (ManagedApplicationResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := applications.ParseApplicationID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ManagedApplication.ApplicationClient.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %v", *id, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (r ManagedApplicationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application" "test" {
  name                        = "acctestManagedApp%d"
  location                    = azurerm_resource_group.test.location
  resource_group_name         = azurerm_resource_group.test.name
  kind                        = "ServiceCatalog"
  managed_resource_group_name = "infraGroup%d"
  application_definition_id   = azurerm_managed_application_definition.test.id

  parameters = {
    location                 = azurerm_resource_group.test.location
    storageAccountNamePrefix = "store%s"
    storageAccountType       = "Standard_LRS"
  }
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomString)
}

func (r ManagedApplicationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application" "import" {
  name                        = azurerm_managed_application.test.name
  location                    = azurerm_managed_application.test.location
  resource_group_name         = azurerm_managed_application.test.resource_group_name
  kind                        = "ServiceCatalog"
  managed_resource_group_name = "infraGroup%d"
}
`, r.basic(data), data.RandomInteger)
}

func (r ManagedApplicationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-mapp-%d"
  location = "%s"
}

resource "azurerm_marketplace_agreement" "test" {
  publisher = "cisco"
  offer     = "cisco-meraki-vmx"
  plan      = "cisco-meraki-vmx"
}

resource "azurerm_managed_application" "test" {
  name                        = "acctestCompleteManagedApp%d"
  location                    = azurerm_resource_group.test.location
  resource_group_name         = azurerm_resource_group.test.name
  kind                        = "MarketPlace"
  managed_resource_group_name = "completeInfraGroup%d"

  plan {
    name      = azurerm_marketplace_agreement.test.plan
    product   = azurerm_marketplace_agreement.test.offer
    publisher = azurerm_marketplace_agreement.test.publisher
    version   = "15.37.1"
  }

  parameters = {
    zone                        = "0"
    location                    = azurerm_resource_group.test.location
    merakiAuthToken             = "f451adfb-d00b-4612-8799-b29294217d4a"
    subnetAddressPrefix         = "10.0.0.0/24"
    subnetName                  = "acctestSubnet"
    virtualMachineSize          = "Standard_DS12_v2"
    virtualNetworkAddressPrefix = "10.0.0.0/16"
    virtualNetworkName          = "acctestVnet"
    virtualNetworkNewOrExisting = "new"
    virtualNetworkResourceGroup = "acctestVnetRg"
    vmName                      = "acctestVM"
  }

  tags = {
    ENV = "Test"
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger)
}

func (r ManagedApplicationResource) parameterValues(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_managed_application" "test" {
  name                        = "acctestManagedApp%d"
  location                    = azurerm_resource_group.test.location
  resource_group_name         = azurerm_resource_group.test.name
  kind                        = "ServiceCatalog"
  managed_resource_group_name = "infraGroup%d"
  application_definition_id   = azurerm_managed_application_definition.test.id

  parameter_values = <<VALUES
	{
        "location": {"value": "${azurerm_resource_group.test.location}"},
        "storageAccountNamePrefix": {"value": "store%s"},
        "storageAccountType": {"value": "Standard_LRS"}
	}
  VALUES
}
`, r.template(data), data.RandomInteger, data.RandomInteger, data.RandomString)
}

func (ManagedApplicationResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

data "azurerm_client_config" "test" {}

data "azurerm_role_definition" "test" {
  name = "Contributor"
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-mapp-%d"
  location = "%s"
}

resource "azurerm_managed_application_definition" "test" {
  name                = "acctestManagedAppDef%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  lock_level          = "ReadOnly"
  package_file_uri    = "https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip"
  display_name        = "TestManagedAppDefinition"
  description         = "Test Managed App Definition"
  package_enabled     = true

  authorization {
    service_principal_id = data.azurerm_client_config.test.object_id
    role_definition_id   = split("/", data.azurerm_role_definition.test.id)[length(split("/", data.azurerm_role_definition.test.id)) - 1]
  }
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
