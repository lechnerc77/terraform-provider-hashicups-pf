package hashicups

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type hashicupsProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}
