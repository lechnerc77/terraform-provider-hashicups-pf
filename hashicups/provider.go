package hashicups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &hashicupsProvider{}
}

var _ provider.Provider = (*hashicupsProvider)(nil)

// hashicupsProvider is the provider implementation.
type hashicupsProvider struct{}

func (p *hashicupsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "http"
}

func (p *hashicupsProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *hashicupsProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *hashicupsProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *hashicupsProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}
