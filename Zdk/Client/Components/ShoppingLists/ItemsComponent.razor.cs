using Microsoft.AspNetCore.Components;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ItemsComponent
    {
        [Parameter]
        public ShoppingList? ShoppingList { get; set; }

        public Item NewItem { get; set; } = new Item();

        private IList<Item> items => ShoppingList?.Items?.ToList() ?? new List<Item>();

        private async Task ToggleSoldOut(Item item)
        {
            item.SoldOut = !item.SoldOut;
            await ShoppingListsRepo.UpdateItem(item);
        }

        private async Task OnSubmit()
        {
            await ShoppingListsRepo.SendItem(NewItem);

            NewItem = new Item()
            {
                Amount = 1,
            };
        }
    }
}
