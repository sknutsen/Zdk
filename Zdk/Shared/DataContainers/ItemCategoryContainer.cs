using Zdk.Shared.Models;

namespace Zdk.Shared.DataContainers;

public class ItemCategoryContainer : ContainerBase, IContainer
{
    public int ItemCategoryId { get; set; }
    public int ItemId { get; set; }
    public int CategoryId { get; set; }

    public Item Item { get; set; }
    public Category Category { get; set; }

    public ItemCategoryContainer() { }

    public ItemCategoryContainer(IEntityClass entity)
    {
        Fill(entity);
    }

    public IEntityClass ToEntityClass()
    {
        return new ItemCategory
        {
            ItemCategoryId = ItemCategoryId,
            ItemId = ItemId,
            CategoryId = CategoryId,
            Item = Item,
            Category = Category
        };
    }

    public void Fill(IEntityClass entity)
    {
        ItemCategory itemCategory = (ItemCategory)entity;

        ItemCategoryId = itemCategory.ItemCategoryId;
        ItemId = itemCategory.ItemId;
        CategoryId = itemCategory.CategoryId;
        Item = itemCategory.Item;
        Category = itemCategory.Category;
    }
}
