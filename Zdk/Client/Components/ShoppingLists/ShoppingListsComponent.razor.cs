using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Authorization;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ShoppingListsComponent
    {
        [Parameter]
        public IList<ShoppingList> ShoppingLists { get; set; } = new List<ShoppingList>();

        [Parameter]
        public ShoppingList? ShoppingList { get; set; }

        [Parameter]
        public EventCallback<ShoppingList?> ShoppingListChanged { get; set; }

        private ShoppingList? shoppingList { get => ShoppingList; set => ShoppingListChanged.InvokeAsync(value); }

        public ShoppingList NewShoppingList { get; set; } = new();

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            AuthenticationState state = await (_currentAuthenticationStateTask);
        }

        private async Task OnSubmit()
        {
            NewShoppingList.Items = new List<ShoppingListItem>();
            await ShoppingListsRepo.SendList(NewShoppingList);

            NewShoppingList = new ShoppingList();
        }
    }
}