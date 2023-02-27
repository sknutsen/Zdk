using System.Collections.Generic;
using Zdk.Shared.Models;

namespace Zdk.Shared.DataContainers;

public class CategoryContainer : ContainerBase, IContainer
{
    public int CategoryId { get; set; }
    public string CategoryName { get; set; }

    public ICollection<ItemCategory> ItemCategories { get; set; }

    public CategoryContainer() { }

    public CategoryContainer(IEntityClass entity)
    {
        Fill(entity);
    }

    public IEntityClass ToEntityClass()
    {
        return new Category
        {
            CategoryId = CategoryId,
            CategoryName = CategoryName,
            ItemCategories = ItemCategories
        };
    }

    public void Fill(IEntityClass entity)
    {
        Category category = (Category)entity;

        CategoryId = category.CategoryId;
        CategoryName = category.CategoryName;
        ItemCategories = category.ItemCategories;
    }
}
