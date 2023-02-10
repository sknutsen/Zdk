using Microsoft.AspNetCore.Components;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ItemsComponent
    {
        [Parameter]
        public ShoppingList? ShoppingList { get; set; }

        [Parameter]
        public EventCallback<ShoppingList?> ShoppingListChanged { get; set; }

        private ShoppingList? shoppingList { get => ShoppingList; set => ShoppingListChanged.InvokeAsync(value); }

        [Parameter]
        public Item? NewItem { get; set; }

        [Parameter]
        public EventCallback<Item> NewItemChanged { get; set; }

        private Item newItem { get => NewItem; set => NewItemChanged.InvokeAsync(value); }

        [Parameter]
        public Func<Item, Task>? Update { get; set; }

        [Parameter]
        public Func<Item, Task>? Delete { get; set; }

        [Parameter]
        public EventCallback<EventArgs> OnSubmit { get; set; }

        private IReadOnlyList<Item> items => ShoppingList?.Items?.ToList() ?? new List<Item>();

        private async Task ToggleSoldOut(Item item)
        {
            item.SoldOut = !item.SoldOut;
            await Update(item);
        }
    }
}
