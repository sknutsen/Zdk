using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    public async Task GetGroupMemberships(int groupId = 0)
    {
        if (groupId == 0)
        {
            groupId = await GetGroupId();

            if (groupId == 0)
            {
                await Clients.Caller.SendAsync(UserManagementHubMethodNames.GetGroupMemberships, new List<GroupMembership>());
                return;
            }
        }

        var groupMemberships = await groupMembershipHandler.List(groupId);
        await Clients.Caller.SendAsync(UserManagementHubMethodNames.GetGroupMemberships, groupMemberships);
    }

    public async Task AddUser(UserContainer user)
    {
        if (!await IsAdmin())
        {
            return;
        }

        int groupId = await GetGroupId();
        await groupMembershipHandler.New(groupId, user.Id, false);
        await GetGroupMemberships();
    }

    public async Task RemoveUser(UserContainer user)
    {
        if (!await IsAdmin())
        {
            return;
        }

        int groupId = await GetGroupId();
        await groupMembershipHandler.Delete(groupId, user.Id);
        await GetGroupMemberships();
    }
}
