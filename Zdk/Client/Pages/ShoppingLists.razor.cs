using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components;
using System.Net;
using System.Net.Http;
using System.Net.Http.Json;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.Components.Forms;
using Microsoft.AspNetCore.Components.Routing;
using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.Web.Virtualization;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.WebAssembly.Http;
using Microsoft.AspNetCore.SignalR.Client;
using Microsoft.JSInterop;
using Zdk.Client;
using Zdk.Shared;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class ShoppingLists : AuthenticatedBase, IAsyncDisposable
    {
        private IList<Group>? _groups;
        private IList<ShoppingList> _shoppingLists = new List<ShoppingList>();
        private ShoppingList? _shoppingList;
        private UserSession? _session;
        private int SessionGroup
        {
            get => _session.GroupId;
            set
            {
                _session.GroupId = value;
            }
        }

        public override async Task SetParametersAsync(ParameterView parameters)
        {
            await base.SetParametersAsync(parameters);
            if (ShoppingListsRepo.IsConnected)
            {
                await ShoppingListsRepo.GetLists();
            }
        }

        protected override async Task OnInitializedAsync()
        {
            await base.OnInitializedAsync();
            try
            {
                await ShoppingListsRepo.Init();
                ShoppingListsRepo.OnReceiveData += ShoppingListsRepo_OnReceiveData;

                await UserRepo.Init();
                UserRepo.OnGetGroups += UserRepo_OnGetGroups;
                UserRepo.OnGetUserSession += UserRepo_OnGetUserSession;
            }
            catch (AccessTokenNotAvailableException exception)
            {
                exception.Redirect();
            }
            catch (HttpRequestException exception)
            {
                if (exception.StatusCode == HttpStatusCode.Unauthorized)
                {
                // NavigationManager.NavigateTo($"authentication/login?returnUrl={Uri.EscapeDataString(NavigationManager.Uri)}");
                }
            }
        }

        private async Task ShoppingListsRepo_OnReceiveData(IList<ShoppingList> list)
        {
            _shoppingLists = list;
            StateHasChanged();
            await Task.CompletedTask;
        }

        private async Task UserRepo_OnGetGroups(IList<Group> list)
        {
            _groups = list;
            Console.WriteLine($"Received groups");
            StateHasChanged();
            await Task.CompletedTask;
        }

        private async Task UserRepo_OnGetUserSession(UserSession session)
        {
            _session = session;
            await JoinClientGroup();
            await ShoppingListsRepo.GetLists();
            StateHasChanged();
        }

        private async Task ChangeGroup(object? value)
        {
            Console.WriteLine($"value is: {value} - type: {value.GetType()}");
            await LeaveClientGroup();
            _session.GroupId = Int32.Parse((string)value);
            await UserRepo.UpdateUserSession(_session);
        }

        private async Task JoinClientGroup()
        {
            if (IsAuthenticated())
            {
                await ShoppingListsRepo.JoinClientGroup();
                await UserRepo.JoinClientGroup();
            }
        }

        private async Task LeaveClientGroup()
        {
            if (IsAuthenticated())
            {
                await ShoppingListsRepo.LeaveClientGroup(SessionGroup);
                await UserRepo.LeaveClientGroup(SessionGroup);
            }
        }

        public async ValueTask DisposeAsync()
        {
            ShoppingListsRepo.OnReceiveData -= ShoppingListsRepo_OnReceiveData;
            UserRepo.OnGetGroups -= UserRepo_OnGetGroups;
            UserRepo.OnGetUserSession -= UserRepo_OnGetUserSession;

            await ShoppingListsRepo.Disconnect();
            await UserRepo.Disconnect();
            await Task.CompletedTask;
        }
    }
}