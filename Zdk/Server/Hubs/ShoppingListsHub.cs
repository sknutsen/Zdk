using Microsoft.AspNetCore.Authorization;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class ShoppingListsHub : BaseHub<ShoppingListsHub>
{
    private readonly ItemHandler itemHandler;
    private readonly ShoppingListHandler shoppingListHandler;

    public ShoppingListsHub(IServiceProvider services) : base(services)
    {
        itemHandler = services.GetRequiredService<ItemHandler>();
        shoppingListHandler = services.GetRequiredService<ShoppingListHandler>();
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
