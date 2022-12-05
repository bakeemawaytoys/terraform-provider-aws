// Code generated by internal/generate/servicepackagedata/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
)

var spd = &servicePackageData{}

func registerFrameworkDataSourceFactory(factory func(context.Context) (datasource.DataSourceWithConfigure, error)) {
	spd.frameworkDataSourceFactories = append(spd.frameworkDataSourceFactories, factory)
}

func registerFrameworkResourceFactory(factory func(context.Context) (resource.ResourceWithConfigure, error)) {
	spd.frameworkResourceFactories = append(spd.frameworkResourceFactories, factory)
}

type servicePackageData struct {
	frameworkDataSourceFactories []func(context.Context) (datasource.DataSourceWithConfigure, error)
	frameworkResourceFactories   []func(context.Context) (resource.ResourceWithConfigure, error)
}

func (d *servicePackageData) Configure(ctx context.Context, meta any) error {
	return nil
}

func (d *servicePackageData) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return d.frameworkDataSourceFactories
}

func (d *servicePackageData) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return d.frameworkResourceFactories
}

func (d *servicePackageData) ServicePackageName() string {
	return "quicksight"
}

var ServicePackageData intf.ServicePackageData = spd
