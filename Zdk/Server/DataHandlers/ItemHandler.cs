using Microsoft.EntityFrameworkCore;
using Zdk.Server.Data;
using Zdk.Server.DataHandlers.Base;
using Zdk.Shared.Models;

namespace Zdk.Server.DataHandlers;

public class ItemHandler : HandlerBase<ItemHandler, DataContext>
{
    public ItemHandler(DataContext context, ILogger<ItemHandler> logger) 
        : base(context, logger)
    {
    }

    public async Task<IList<Item>> List(int shoppingListId)
    {
        return await context.Items.Where(e => e.ShoppingListId == shoppingListId).ToListAsync();
    }

    public async Task<Item?> Get(int id)
    {
        return await context.Items.FindAsync(id);
    }

    public async Task<bool> New(Item item)
    {
        try
        {
            item.CreatedAt = DateTime.Now;
            item.UpdatedAt = DateTime.Now;
            await context.Items.AddAsync(item);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Update(Item item)
    {
        try
        {
            item.UpdatedAt = DateTime.Now;
            context.Items.Update(item);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Delete(Item item)
    {
        try
        {
            context.Items.Remove(item);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }
}
