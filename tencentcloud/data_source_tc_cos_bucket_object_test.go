package tencentcloud

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCosBucketObjectDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCosBucketObjectDataSource(appid),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCosBucketObjectExists("cloud_cos_bucket_object.object_content"),
					resource.TestCheckResourceAttr("data.cloud_cos_bucket_object.object", "content_type", "binary/octet-stream"),
					resource.TestMatchResourceAttr("data.cloud_cos_bucket_object.object", "last_modified",
						regexp.MustCompile("^[a-zA-Z]{3}, [0-9]+ [a-zA-Z]+ [0-9]{4} [0-9:]+ [A-Z]+$")),
				),
			},
		},
	})
}

func testAccCosBucketObjectDataSource(appid string) string {
	return fmt.Sprintf(`
resource "cloud_cos_bucket" "object_bucket" {
  bucket = "tf-bucket-%d-%s"
}

resource "cloud_cos_bucket_object" "object_content" {
  bucket       = cloud_cos_bucket.object_bucket.bucket
  key          = "tf-object-content"
  content      = "aaaaaaaaaaaaaaaa"
  content_type = "binary/octet-stream"
}

data "cloud_cos_bucket_object" "object" {
  bucket = cloud_cos_bucket_object.object_content.bucket
  key    = cloud_cos_bucket_object.object_content.key
}
`, acctest.RandInt(), appid)
}
