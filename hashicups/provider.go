package hashicups

import (
	"context"
	"os"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &hashicupsProvider{}
}

var _ provider.Provider = (*hashicupsProvider)(nil)

// hashicupsProvider is the provider implementation.
type hashicupsProvider struct{}

func (p *hashicupsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hashicups"
}

func (p *hashicupsProvider) Schema(_ context.Context, req provider.SchemaRequest, res *provider.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:            true,
				Validators:          []validator.String{},
				MarkdownDescription: "Hostname",
			},
			"username": schema.StringAttribute{
				Optional:            true,
				Validators:          []validator.String{},
				MarkdownDescription: "Username",
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Validators:          []validator.String{},
				MarkdownDescription: "Password",
				Sensitive:           true,
			},
		},
		MarkdownDescription: "HashiCups Provider Schema",
	}
}

func (p *hashicupsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring HashiCups client")

	// Load configuration into model
	var config hashicupsProviderModel
	diag := req.Config.Get(ctx, &config)

	res.Diagnostics.Append(diag...)
	if res.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown HashiCups API Host",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HASHICUPS_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown HashiCups API Username",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HASHICUPS_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown HashiCups API Password",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the HASHICUPS_PASSWORD environment variable.",
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.
	host := os.Getenv("HASHICUPS_HOST")
	username := os.Getenv("HASHICUPS_USERNAME")
	password := os.Getenv("HASHICUPS_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if host == "" {
		res.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing HashiCups API Host",
			"The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API host. "+
				"Set the host value in the configuration or use the HASHICUPS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		res.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing HashiCups API Username",
			"The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API username. "+
				"Set the username value in the configuration or use the HASHICUPS_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		res.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing HashiCups API Password",
			"The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API password. "+
				"Set the password value in the configuration or use the HASHICUPS_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if res.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "hashicups_host", host)
	ctx = tflog.SetField(ctx, "hashicups_username", username)
	ctx = tflog.SetField(ctx, "hashicups_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "hashicups_password")

	tflog.Debug(ctx, "Creating HashiCups client")

	// Create a new HashiCups client using the configuration values
	client, err := hashicups.NewClient(&host, &username, &password)
	if err != nil {
		res.Diagnostics.AddError(
			"Unable to Create HashiCups API Client",
			"An unexpected error occurred when creating the HashiCups API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"HashiCups Client Error: "+err.Error(),
		)
		return
	}
	// Make the HashiCups client available during DataSource and Resource
	// type Configure methods.
	res.DataSourceData = client
	res.ResourceData = client

	tflog.Info(ctx, "Configured HashiCups client", map[string]any{"success": true})
}

func (p *hashicupsProvider) Resources(ctx context.Context) []func() resource.Resource {

	tflog.Warn(ctx, "Resources data fetched")
	return []func() resource.Resource{
		NewOrderResource,
	}
}

func (p *hashicupsProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCoffeesDataSource,
	}
}
