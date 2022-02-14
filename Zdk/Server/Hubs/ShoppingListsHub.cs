﻿using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared;
using Microsoft.AspNetCore.Authorization;
using Zdk.Shared.Models;
using Zdk.Utilities.Authentication.Helpers;
using Zdk.Client.Pages;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

[Authorize]
public class ShoppingListsHub : Hub
{
    private readonly ItemHandler itemHandler;
    private readonly ShoppingListHandler shoppingListHandler;
    private readonly GroupHandler groupHandler;
    private readonly GroupMembershipHandler groupMembershipHandler;
    private readonly UserSessionHandler userSessionHandler;
    private readonly ILogger<ShoppingListsHub> logger;

    public ShoppingListsHub(ItemHandler itemHandler, ShoppingListHandler shoppingListHandler, GroupHandler groupHandler, GroupMembershipHandler groupMembershipHandler, UserSessionHandler userSessionHandler, ILogger<ShoppingListsHub> logger)
    {
        this.itemHandler = itemHandler;
        this.shoppingListHandler = shoppingListHandler;
        this.groupHandler = groupHandler;
        this.groupMembershipHandler = groupMembershipHandler;
        this.userSessionHandler = userSessionHandler;
        this.logger = logger;
    }

    public async Task ListShoppingLists()
    {
        await JoinGroup();

        int groupId = await GetGroupId();
        var lists = await shoppingListHandler.List(groupId);
        await Clients.Group(groupId.ToString()).SendAsync(ShoppingListsHubMethodNames.ReceiveData, lists);
    }

    public async Task SendShoppingList(ShoppingList shoppingList)
    {
        shoppingList.PostedBy = Context.User?.GetUserId() ?? "???";
        shoppingList.GroupId = await GetGroupId();

        await shoppingListHandler.New(shoppingList);

        await ListShoppingLists();
    }

    public async Task SendItem(Item item)
    {
        if (item.PostedBy == null)
        {
            item.PostedBy = Context.User?.GetUserId() ?? "???";

            if (item.PostedBy == "???")
            {
                logger.LogWarning("Item posted without userId");
            }
        }

        await itemHandler.New(item);

        logger.LogInformation($"Item posted with userid: {item.PostedBy}");

        await ListShoppingLists();
    }

    public async Task UpdateShoppingList(ShoppingList shoppingList)
    {
        shoppingList.PostedBy = Context.User?.GetUserId() ?? "???";

        await shoppingListHandler.Update(shoppingList);

        await ListShoppingLists();
    }

    public async Task UpdateItem(Item item)
    {
        item.PostedBy = Context.User?.GetUserId() ?? "???";

        await itemHandler.Update(item);

        await ListShoppingLists();
    }

    public async Task DeleteShoppingList(ShoppingList shoppingList)
    {
        await shoppingListHandler.Delete(shoppingList);

        await ListShoppingLists();
    }

    public async Task DeleteItem(Item item)
    {
        await itemHandler.Delete(item);

        await ListShoppingLists();
    }

    public override async Task OnConnectedAsync()
    {
        await ListShoppingLists();

        await base.OnConnectedAsync();
    }

    public override async Task OnDisconnectedAsync(Exception? exception)
    {
        await LeaveGroup();

        await base.OnDisconnectedAsync(exception);
    }

    private async Task<int> GetGroupId()
    {
        string? userId = Context.User?.GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        return session.GroupId;
    }

    public async Task JoinGroup()
    {
        int groupId = await GetGroupId();

        await JoinGroup(groupId);
    }

    private async Task JoinGroup(int groupId)
    {
        await Groups.AddToGroupAsync(Context.ConnectionId, groupId.ToString());
    }

    private async Task LeaveGroup()
    {
        int groupId = await GetGroupId();

        await LeaveGroup(groupId);
    }

    public async Task LeaveGroup(int groupId)
    {
        await Groups.RemoveFromGroupAsync(Context.ConnectionId, groupId.ToString());
    }
}
