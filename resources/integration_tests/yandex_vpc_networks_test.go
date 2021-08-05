package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/yandex-cloud/cq-provider-yandex/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationVPCNetworks(t *testing.T) {
	var tfTmpl = fmt.Sprintf(`
resource "yandex_vpc_network" "cq-net-test-net-%[1]s" {
  name = "cq-net-test-net-%[1]s"
}
`, suffix)
	testIntegrationHelper(t, resources.VPCNetworks(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "yandex_vpc_networks",
			Filter: func(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("cq-net-test-net-%s", suffix),
				},
			}},
		}
	}, tfTmpl)
}
