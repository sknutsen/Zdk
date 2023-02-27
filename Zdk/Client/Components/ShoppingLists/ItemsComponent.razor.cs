using Microsoft.AspNetCore.Components;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ItemsComponent
    {
        [Parameter]
        public ShoppingList? ShoppingList { get; set; }

        public ShoppingListItem NewItem { get; set; } = new ShoppingListItem();

        private IList<ShoppingListItem> items => ShoppingList?.Items?.ToList() ?? new List<ShoppingListItem>();

        private async Task ToggleSoldOut(ShoppingListItem item)
        {
            item.SoldOut = !item.SoldOut;
            await ShoppingListsRepo.UpdateItem(item);
        }

        private async Task OnSubmit()
        {
            await ShoppingListsRepo.SendItem(NewItem);

            NewItem = new ShoppingListItem()
            {
                Amount = 1,
            };
        }
    }
}
