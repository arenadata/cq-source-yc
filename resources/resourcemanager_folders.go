package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManagerFolders() *schema.Table {
	return &schema.Table{
		Name:        "yandex_resourcemanager_folders",
		Resolver:    fetchResourceManagerFolders,
		Multiplex:   client.MultiplexBy(client.Folders),
		IgnoreError: client.IgnoreErrorHandler,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"cloud_id", "id"}},
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the folder.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, Unique: true},
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the folder. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the folder. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "cloud_id",
				Type:        schema.TypeString,
				Description: "ID of the organization that the folder belongs to.",
				Resolver:    schema.PathResolver("CloudId"),
			},
		},
	}

}

func fetchResourceManagerFolders(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	folder, err := c.Services.ResourceManager.Folder().Get(ctx, &resourcemanager.GetFolderRequest{FolderId: c.MultiplexedResourceId})
	if err != nil {
		return err
	}

	res <- folder

	return nil
}
