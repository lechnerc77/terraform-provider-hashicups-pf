package hashicups

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type coffeesDataSource struct {
	client *hashicups.Client
}

var _ datasource.DataSource = (*coffeesDataSource)(nil)
var _ datasource.DataSourceWithConfigure = (*coffeesDataSource)(nil)

func NewCoffeesDataSource() datasource.DataSource {
	return &coffeesDataSource{}
}

func (d *coffeesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coffees"
}

func (d *coffeesDataSource) Schema(_ context.Context, req datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Description:         "Fetches the list of coffees.",
		MarkdownDescription: "Fetches the list of coffees.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Placeholder identifier attribute.",
				MarkdownDescription: "Placeholder identifier attribute.",
			},
			"coffees": schema.ListNestedAttribute{
				Description:         "List of coffees.",
				MarkdownDescription: "List of coffees.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							Description:         "Numeric identifier of the coffee.",
							MarkdownDescription: "Numeric identifier of the coffee.",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							Description:         "Product name of the coffee.",
							MarkdownDescription: "Product name of the coffee.",
						},
						"teaser": schema.StringAttribute{
							Computed:            true,
							Description:         "Fun tagline for the coffee.",
							MarkdownDescription: "Fun tagline for the coffee.",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							Description:         "Product description of the coffee.",
							MarkdownDescription: "Product description of the coffee.",
						},
						"price": schema.Float64Attribute{
							Computed:            true,
							Description:         "Suggested cost of the coffee.",
							MarkdownDescription: "Suggested cost of the coffee.",
						},
						"image": schema.StringAttribute{
							Computed:            true,
							Description:         "URI for an image of the coffee.",
							MarkdownDescription: "URI for an image of the coffee.",
						},
						"ingredients": schema.ListNestedAttribute{
							Description:         "List of ingredients in the coffee.",
							MarkdownDescription: "List of ingredients in the coffee.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Computed:            true,
										Description:         "Numeric identifier of the coffee ingredient.",
										MarkdownDescription: "Numeric identifier of the coffee ingredient.",
									},
								},
							},
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (d *coffeesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state coffeesDataSourceModel

	coffees, err := d.client.GetCoffees()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HashiCups Coffees",
			err.Error(),
		)
		return
	}

	for _, coffee := range coffees {
		coffeeState := coffeesModel{
			ID:          types.Int64Value(int64(coffee.ID)),
			Name:        types.StringValue(coffee.Name),
			Teaser:      types.StringValue(coffee.Teaser),
			Description: types.StringValue(coffee.Description),
			Price:       types.Float64Value(coffee.Price),
			Image:       types.StringValue(coffee.Image),
		}
		for _, ingredient := range coffee.Ingredient {
			coffeeState.Ingredients = append(coffeeState.Ingredients, coffeesIngredientsModel{
				ID: types.Int64Value(int64(ingredient.ID)),
			})
		}
		state.Coffees = append(state.Coffees, coffeeState)
	}

	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *coffeesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*hashicups.Client)
}
