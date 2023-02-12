using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class ShoppingListsHub : BaseHub<ShoppingListsHub>
{
    public async Task ListShoppingLists()
    {
        int groupId = await GetGroupId();

        await JoinGroup(groupId);

        await ListShoppingListsByGroup(groupId);
    }

    public async Task ListShoppingListsByGroup(int groupId)
    {
        var lists = await shoppingListHandler.List(groupId);
        await Clients.Group(groupId.ToString()).SendAsync(ShoppingListsHubMethodNames.ReceiveData, lists);
    }

    public async Task SendShoppingList(ShoppingList shoppingList)
    {
        string userId = GetUserId();
        int groupId = await GetGroupId();

        shoppingList.PostedBy = userId;
        shoppingList.GroupId = groupId;

        ShoppingList? addedList = await shoppingListHandler.New(shoppingList);

        if (addedList is null)
        {
            logger.LogWarning("failed adding new list");
            await ListShoppingLists();
            return;
        }

        IList<ShoppingList> shoppingLists = await shoppingListHandler.List(addedList.GroupId);

        await itemHandler.MoveSoldOutItems(shoppingLists, addedList.ShoppingListId);

        await ListShoppingListsByGroup(groupId);
    }

    public async Task UpdateShoppingList(ShoppingList shoppingList)
    {
        shoppingList.PostedBy = GetUserId();

        await shoppingListHandler.Update(shoppingList);

        await ListShoppingLists();
    }

    public async Task DeleteShoppingList(ShoppingList shoppingList)
    {
        await shoppingListHandler.Delete(shoppingList);

        await ListShoppingLists();
    }
}
