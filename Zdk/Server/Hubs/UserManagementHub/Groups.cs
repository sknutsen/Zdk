using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    public async Task GetGroups()
    {
        string userId = GetUserId();

        if (string.IsNullOrEmpty(userId))
        {
            return;
        }

        UserSession session = await userSessionHandler.Get(userId);

        var groups = await groupHandler.List(userId, session.IsAdmin);

        int groupId = await GetGroupId();

        await Clients.Group(groupId.ToString()).SendAsync(UserManagementHubMethodNames.GetGroups, groups);
    }

    public async Task NewGroup(Group group)
    {
        string userId = GetUserId();

        if (string.IsNullOrEmpty(group.Owner))
        {
            group.Owner = userId;
        }

        (bool created, Group? newGroup) = await groupHandler.New(group);

        if (!created || newGroup is null)
        {
            await GetGroups();
            return;
        }

        await groupMembershipHandler.New(newGroup.GroupId, userId, false);

        await GetGroups();
    }

    public async Task UpdateGroup(Group group)
    {
        await groupHandler.Update(group);

        await GetGroups();
    }

    public async Task DeleteGroup(Group group)
    {
        await groupHandler.Delete(group);

        await GetGroups();
    }
}
