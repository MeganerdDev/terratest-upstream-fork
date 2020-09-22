// +build azure azureslim,network

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods can be mocked or Create/Delete APIs are added, these tests can be extended.
*/

func TestGetVirtualNetworkE(t *testing.T) {
	t.Parallel()

	vnetName := ""
	rgName := ""
	subID := ""

	_, err := GetVirtualNetworkE(t, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestGetSubnetE(t *testing.T) {
	t.Parallel()

	subnetName := ""
	vnetName := ""
	rgName := ""
	subID := ""

	_, err := GetSubnetE(t, subnetName, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestGetVirtualNetworkDNSServerIPsE(t *testing.T) {
	t.Parallel()

	vnetName := ""
	rgName := ""
	subID := ""

	_, err := GetVirtualNetworkDNSServerIPsE(t, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestGetVirtualNetworkSubnetsE(t *testing.T) {
	t.Parallel()

	vnetName := ""
	rgName := ""
	subID := ""

	_, err := GetVirtualNetworkSubnetsE(t, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestCheckSubnetContainsIPE(t *testing.T) {
	t.Parallel()

	ipAddress := ""
	subnetName := ""
	vnetName := ""
	rgName := ""
	subID := ""

	_, err := CheckSubnetContainsIPE(t, ipAddress, subnetName, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestSubnetExistsE(t *testing.T) {
	t.Parallel()

	subnetName := ""
	vnetName := ""
	rgName := ""
	subID := ""

	_, err := SubnetExistsE(t, subnetName, vnetName, rgName, subID)

	require.Error(t, err)
}

func TestVirtualNetworkExistsE(t *testing.T) {
	t.Parallel()

	vnetName := ""
	rgName := ""
	subID := ""

	_, err := VirtualNetworkExistsE(t, vnetName, rgName, subID)

	require.Error(t, err)
}
