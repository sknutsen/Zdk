using Microsoft.AspNetCore.Authorization;
using Microsoft.Extensions.Logging;
using Zdk.Shared.Models;

namespace Zdk.Server.ShoppingLists;

[Authorize]
public partial class ShoppingListsHub : SlHubBase<ShoppingListsHub>
{
    public async Task SendItem(ShoppingListItem item)
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

    public async Task UpdateItem(ShoppingListItem item)
    {
        item.PostedBy = GetUserId();

        await itemHandler.Update(item);

        await ListShoppingLists();
    }

    public async Task DeleteItem(ShoppingListItem item)
    {
        await itemHandler.Delete(item);

        await ListShoppingLists();
    }
}
