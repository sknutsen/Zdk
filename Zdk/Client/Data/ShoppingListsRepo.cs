using Microsoft.AspNetCore.SignalR.Client;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Client;

public class ShoppingListsRepo : RepoBase, IAsyncDisposable
{
    private List<ShoppingList> listsToSend = new List<ShoppingList>();
    private List<Item> itemsToSend = new List<Item>();

    public event Func<IList<ShoppingList>, Task> OnReceiveData;

    public ShoppingListsRepo(IServiceProvider services) : base(services)
    {
    }

    protected override string GetHubName() => ShoppingListsHubMethodNames.HubName;

    protected override async Task HubConnectionOnReconnected(string? arg)
    {
        await SendLists();
        await SendItems();
    }

    protected override async Task<IList<IDisposable>> On()
    {
        await Task.CompletedTask;

        IDisposable receiveData = HubConnection!.On<List<ShoppingList>>(ShoppingListsHubMethodNames.ReceiveData, async (list) =>
        {
            await OnReceiveData?.Invoke(list);
        });

        return new List<IDisposable>()
        {
            receiveData,
        };
    }

    public async Task GetLists()
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(ShoppingListsHubMethodNames.ListShoppingLists);
            Console.WriteLine($"Getting lists");
        }
    }

    public async Task SendList(ShoppingList newShoppingList)
    {
        listsToSend.Add(newShoppingList);

        await SendLists();
    }

    private async Task SendLists()
    {
        if (IsConnected)
        {
            foreach (var list in listsToSend)
            {
                await HubConnection!.SendAsync(ShoppingListsHubMethodNames.SendShoppingList, list);
            }

            listsToSend.Clear();
        }
    }

    public async Task SendItem(Item newItem)
    {
        itemsToSend.Add(newItem);

        await SendItems();
    }

    private async Task SendItems()
    {
        if (IsConnected)
        {
            foreach (var item in itemsToSend)
            {
                await HubConnection!.SendAsync(ShoppingListsHubMethodNames.SendItem, item);
            }
        }
    }

    public async Task UpdateList(ShoppingList list)
    {
        await HubConnection!.SendAsync(ShoppingListsHubMethodNames.UpdateShoppingList, list);
    }

    public async Task UpdateItem(Item item)
    {
        await HubConnection!.SendAsync(ShoppingListsHubMethodNames.UpdateItem, item);
    }

    public async Task DeleteList(ShoppingList list)
    {
        await HubConnection!.SendAsync(ShoppingListsHubMethodNames.DeleteShoppingList, list);
    }

    public async Task DeleteItem(Item item)
    {
        await HubConnection!.SendAsync(ShoppingListsHubMethodNames.DeleteItem, item);
    }
}
