using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using Zdk.Shared.Models;

namespace Zdk.DataAccess;

public class ItemHandler : HandlerBase<ItemHandler, ShoppingListContext>
{
    public ItemHandler(ShoppingListContext context, ILogger<ItemHandler> logger) 
        : base(context, logger)
    {
    }

    public async Task<IList<ShoppingListItem>> List(int shoppingListId)
    {
        return await context.Items.Where(e => e.ShoppingListId == shoppingListId).ToListAsync();
    }

    public async Task<IList<ShoppingListItem>> List()
    {
        return await context.Items.ToListAsync();
    }

    public async Task<ShoppingListItem?> Get(int id)
    {
        return await context.Items.FindAsync(id);
    }

    public async Task<ShoppingListItem?> New(ShoppingListItem item)
    {
        try
        {
            item.CreatedAt = DateTime.Now;
            item.UpdatedAt = DateTime.Now;
            var addResult = await context.Items.AddAsync(item);
            await context.SaveChangesAsync();

            return addResult.Entity;
        }
        catch (Exception e)
        {
            logger.LogError(e, "New item failed");

            return null;
        }
    }

    public async Task<bool> Update(ShoppingListItem item)
    {
        try
        {
            item.UpdatedAt = DateTime.Now;
            context.Items.Update(item);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception e)
        {
            logger.LogError(e, $"Update item failed for item with id: {item.ItemId}");

            return false;
        }
    }

    public async Task<bool> Delete(ShoppingListItem item)
    {
        try
        {
            context.Items.Remove(item);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception e)
        {
            logger.LogError(e, $"Delete item failed for item with id: {item.ItemId}");

            return false;
        }
    }

    public async Task<bool> MoveSoldOutItems(IList<ShoppingList> shoppingLists, int newList)
    {
        bool result = false;

        try
        {
            foreach (ShoppingList shoppingList in shoppingLists)
            {
                result = await MoveSoldOutItems(shoppingList, newList);
            }

            return result;
        }
        catch (Exception e)
        {
            logger.LogError(e, "MoveSoldOutItems failed");

            return result;
        }
    }

    public async Task<bool> MoveSoldOutItems(ShoppingList shoppingList, int newList)
    {
        bool result = false;

        try
        {
            IList<ShoppingListItem> items = shoppingList.Items.Where(e => e.SoldOut).ToList();
            foreach (ShoppingListItem item in items)
            {
                item.SoldOut = false;
                item.ShoppingListId = newList;

                result = await Update(item);
            }

            return result;
        }
        catch (Exception e)
        {
            logger.LogError(e, $"MoveSoldOutItems failed for old list: {shoppingList.ShoppingListId} -- new: {newList}");

            return result;
        }
    }
}
