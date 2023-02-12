using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    public async Task GetSession()
    {
        string userId = GetUserId();

        if (userId == null)
        {
            throw new Exception();
        }

        var session = await userSessionHandler.Get(userId);

        await Clients.Caller.SendAsync(UserManagementHubMethodNames.GetUserSession, session);
        await GetGroupMemberships();
    }

    public async Task UpdateSession(UserSession session)
    {
        await userSessionHandler.Update(session);
        await JoinGroup();

        await GetSession();
    }
}
