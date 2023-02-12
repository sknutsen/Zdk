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
    [Authorize]
    public partial class UserManagement : AuthenticatedBase, IAsyncDisposable
    {
        private IList<Group> groups = new List<Group>();
        private IList<GroupMembership> groupMemberships = new List<GroupMembership>();
        private IList<UserContainer> users = new List<UserContainer>();
        private UserSession session;
        private int SessionGroup 
        { 
            get => session.GroupId; 
            set => session.GroupId = value; 
        }

        protected override async Task OnInitializedAsync()
        {
            try
            {
                await ShoppingListsRepo.Init();

                await UserRepo.Init();
                UserRepo.OnGetGroups += UserRepo_OnGetGroups;
                UserRepo.OnGetGroupMemberships += UserRepo_OnGetGroupMemberships;
                UserRepo.OnGetUsers += UserRepo_OnGetUsers;
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
                    NavigationManager.NavigateTo($"authentication/login?returnUrl={Uri.EscapeDataString(NavigationManager.Uri)}");
                }
            }
        }

        private async Task UserRepo_OnGetGroups(IList<Group> list)
        {
            groups = list;
            Console.WriteLine($"Received groups");
            StateHasChanged();
            await Task.CompletedTask;
        }

        private async Task UserRepo_OnGetGroupMemberships(IList<GroupMembership> list)
        {
            groupMemberships = list;
            StateHasChanged();
            await Task.CompletedTask;
        }

        private async Task UserRepo_OnGetUsers(IList<UserContainer> list)
        {
            users = list;
            StateHasChanged();
            await Task.CompletedTask;
        }

        private async Task UserRepo_OnGetUserSession(UserSession session)
        {
            this.session = session;
            await JoinClientGroup();
            StateHasChanged();
        }

        private async Task ChangeGroup(object? value)
        {
            Console.WriteLine($"value is: {value} - type: {value.GetType()}");
            await LeaveClientGroup();
            session.GroupId = Int32.Parse((string)value);
            await UserRepo.UpdateUserSession(session);
        }

        private async Task CreateGroup(string groupName)
        {
            Group group = new Group{GroupName = groupName, Owner = session.UserId, };
            await UserRepo.NewGroup(group);
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
            UserRepo.OnGetGroups -= UserRepo_OnGetGroups;
            UserRepo.OnGetGroupMemberships -= UserRepo_OnGetGroupMemberships;
            UserRepo.OnGetUsers -= UserRepo_OnGetUsers;
            UserRepo.OnGetUserSession -= UserRepo_OnGetUserSession;

            await ShoppingListsRepo.Disconnect();
            await UserRepo.Disconnect();
            await Task.CompletedTask;
        }
    }
}