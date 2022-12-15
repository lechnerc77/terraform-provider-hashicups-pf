package hashicups

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type hashicupsProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// coffeesDataSourceModel maps the data source schema data.
type coffeesDataSourceModel struct {
	Coffees []coffeesModel `tfsdk:"coffees"`
}

// coffeesModel maps coffees schema data.
type coffeesModel struct {
	ID          types.Int64               `tfsdk:"id"`
	Name        types.String              `tfsdk:"name"`
	Teaser      types.String              `tfsdk:"teaser"`
	Description types.String              `tfsdk:"description"`
	Price       types.Float64             `tfsdk:"price"`
	Image       types.String              `tfsdk:"image"`
	Ingredients []coffeesIngredientsModel `tfsdk:"ingredients"`
}

// coffeesIngredientsModel maps coffee ingredients data
type coffeesIngredientsModel struct {
	ID types.Int64 `tfsdk:"id"`
}
