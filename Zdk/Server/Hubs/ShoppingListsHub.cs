using Microsoft.AspNetCore.Authorization;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class ShoppingListsHub : BaseHub<ShoppingListsHub>
{
    private readonly ItemHandler itemHandler;
    private readonly ShoppingListHandler shoppingListHandler;

    public ShoppingListsHub(ItemHandler itemHandler, ShoppingListHandler shoppingListHandler, GroupHandler groupHandler, GroupMembershipHandler groupMembershipHandler, UserSessionHandler userSessionHandler, ILogger<ShoppingListsHub> logger)
        : base(groupHandler, groupMembershipHandler, userSessionHandler, logger)
    {
        this.itemHandler = itemHandler;
        this.shoppingListHandler = shoppingListHandler;
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
}
