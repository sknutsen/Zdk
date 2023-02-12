using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    public async Task GetUsers()
    {
        var users = await userHandler.List();

        string userId = GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        if (!session.IsAdmin)
        {
            users = users.Where(e => e.Id == userId).ToList();
        }

        await Clients.Caller.SendAsync(UserManagementHubMethodNames.GetUsers, users.Select(e => new UserContainer { Id = e.Id, Name = e.UserName ?? e.Email }).ToList());
    }
}
