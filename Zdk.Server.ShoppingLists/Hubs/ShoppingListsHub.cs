using Microsoft.AspNetCore.Authorization;
using Microsoft.Extensions.DependencyInjection;
using Zdk.DataAccess;

namespace Zdk.Server.ShoppingLists;

[Authorize]
public partial class ShoppingListsHub : SlHubBase<ShoppingListsHub>
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
