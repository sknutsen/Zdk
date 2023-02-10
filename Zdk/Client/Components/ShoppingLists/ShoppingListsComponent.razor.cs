using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Authorization;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ShoppingListsComponent
    {
        [Parameter]
        public List<ShoppingList> ShoppingLists { get; set; } = new();
        [Parameter]
        public ShoppingList? ShoppingList { get; set; }

        [Parameter]
        public EventCallback<ShoppingList?> ShoppingListChanged { get; set; }

        private ShoppingList? shoppingList { get => ShoppingList; set => ShoppingListChanged.InvokeAsync(value); }

        [Parameter]
        public ShoppingList NewShoppingList { get; set; } = new();
        [Parameter]
        public EventCallback<ShoppingList> NewShoppingListChanged { get; set; }

        private ShoppingList newShoppingList { get => NewShoppingList; set => NewShoppingListChanged.InvokeAsync(value); }

        [Parameter]
        public Func<ShoppingList, Task>? Delete { get; set; }

        [Parameter]
        public Func<Task>? OnSubmit { get; set; }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            AuthenticationState state = await (_currentAuthenticationStateTask);
        }
    }
}