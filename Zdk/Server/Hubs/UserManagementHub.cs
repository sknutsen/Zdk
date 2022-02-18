using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Server.DataHandlers;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;
using Zdk.Utilities.Authentication.Data;
using Zdk.Utilities.Authentication.Helpers;

namespace Zdk.Server.Hubs;

[Authorize]
public class UserManagementHub : Hub
{
    private readonly GroupHandler groupHandler;
    private readonly GroupMembershipHandler groupMembershipHandler;
    private readonly UserHandler userHandler;
    private readonly UserSessionHandler userSessionHandler;
    private readonly ILogger<UserManagementHub> logger;

    public UserManagementHub(GroupHandler groupHandler, GroupMembershipHandler groupMembershipHandler, UserHandler userHandler, UserSessionHandler userSessionHandler, ILogger<UserManagementHub> logger)
    {
        this.groupHandler = groupHandler;
        this.groupMembershipHandler = groupMembershipHandler;
        this.userHandler = userHandler;
        this.userSessionHandler = userSessionHandler;
        this.logger = logger;
    }

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

    public async Task GetGroups()
    {
        string? userId = Context.User?.GetUserId();

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
        string? userId = Context.User?.GetUserId();

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

    public async Task GetUsers()
    {
        var users = await userHandler.List();

        string? userId = Context.User?.GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        if (!session.IsAdmin)
        {
            users = users.Where(e => e.Id == userId).ToList();
        }

        await Clients.Caller.SendAsync(UserManagementHubMethodNames.GetUsers, users.Select(e => new UserContainer { Id = e.Id, Name = e.UserName ?? e.Email}).ToList());
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

    public async Task GetSession()
    {
        string? userId = Context.User?.GetUserId();

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
        //await LeaveGroup();
        await userSessionHandler.Update(session);
        await JoinGroup();

        await GetSession();
        await GetGroupMemberships();
    }

    public override async Task OnConnectedAsync()
    {
        await GetSession();
        await JoinGroup();
        await GetGroups();
        await GetUsers();

        await base.OnConnectedAsync();
    }

    public override async Task OnDisconnectedAsync(Exception? exception)
    {
        await LeaveGroup();

        await base.OnDisconnectedAsync(exception);
    }

    private async Task<int> GetGroupId()
    {
        string? userId = Context.User?.GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        return session.GroupId;
    }

    public async Task JoinGroup()
    {
        int groupId = await GetGroupId();

        await JoinGroup(groupId);
    }

    private async Task JoinGroup(int groupId)
    {
        await Groups.AddToGroupAsync(Context.ConnectionId, groupId.ToString());
    }

    private async Task LeaveGroup()
    {
        int groupId = await GetGroupId();

        await LeaveGroup(groupId);
    }

    public async Task LeaveGroup(int groupId)
    {
        await Groups.RemoveFromGroupAsync(Context.ConnectionId, groupId.ToString());
    }

    private async Task<bool> IsAdmin()
    {
        string? userId = Context.User?.GetUserId();

        return await userSessionHandler.IsAdmin(userId);
    }
}
