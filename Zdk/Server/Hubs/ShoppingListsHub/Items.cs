using Microsoft.AspNetCore.Authorization;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class ShoppingListsHub : BaseHub<ShoppingListsHub>
{
    public async Task SendItem(Item item)
    {
        if (item.PostedBy == null)
        {
            item.PostedBy = GetUserId();

            if (item.PostedBy == "???")
            {
                logger.LogWarning("Item posted without userId");
            }
        }

        await itemHandler.New(item);

        logger.LogInformation($"Item posted with userid: {item.PostedBy}");

        await ListShoppingLists();
    }

    public async Task UpdateItem(Item item)
    {
        item.PostedBy = GetUserId();

        await itemHandler.Update(item);

        await ListShoppingLists();
    }

    public async Task DeleteItem(Item item)
    {
        await itemHandler.Delete(item);

        await ListShoppingLists();
    }
}
