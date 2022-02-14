using Microsoft.EntityFrameworkCore;
using Zdk.Server.Data;
using Zdk.Server.DataHandlers.Base;
using Zdk.Shared.Helpers;
using Zdk.Shared.Models;

namespace Zdk.Server.DataHandlers;

public class ShoppingListHandler : HandlerBase<ShoppingListHandler, DataContext>
{
    public ShoppingListHandler(DataContext context, ILogger<ShoppingListHandler> logger) 
        : base(context, logger)
    {
    }

    public async Task<IList<ShoppingList>> List(int groupId)
    {
        var list = await context.ShoppingLists.Include(e => e.Items).Where(e => e.GroupId == groupId).ToListAsync();
        list.SortAll();

        return list;
    }

    public async Task<ShoppingList?> Get(int id)
    {
        return await context.ShoppingLists.FindAsync(id);
    }

    public async Task<bool> New(ShoppingList shoppingList)
    {
        try
        {
            shoppingList.CreatedAt = DateTime.Now;
            shoppingList.UpdatedAt = DateTime.Now;
            await context.ShoppingLists.AddAsync(shoppingList);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Update(ShoppingList shoppingList)
    {
        try
        {
            shoppingList.UpdatedAt = DateTime.Now;
            context.ShoppingLists.Update(shoppingList);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Delete(ShoppingList shoppingList)
    {
        try
        {
            context.ShoppingLists.Remove(shoppingList);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }
}
