package provider

import (
	"filippo.io/age"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAgeSecretKey,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceCreateKey("age_secret_key.foo"),
				),
			},
		},
	})
}

const testAccAgeSecretKey = `
resource "age_secret_key" "foo" {}
`

func testAccResourceCreateKey(id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("not found: %s", id)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		_, err := age.ParseX25519Identity(rs.Primary.Attributes["secret_key"])
		if err != nil {
			return err
		}

		_, err = age.ParseX25519Recipient(rs.Primary.Attributes["public_key"])
		if err != nil {
			return err
		}

		return nil
	}
}
