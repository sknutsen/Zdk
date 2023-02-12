using Microsoft.AspNetCore.SignalR.Client;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Client;

public class UserRepo : RepoBase, IAsyncDisposable
{
    public event Func<IList<Group>, Task> OnGetGroups;
    public event Func<IList<GroupMembership>, Task> OnGetGroupMemberships;
    public event Func<IList<UserContainer>, Task> OnGetUsers;
    public event Func<UserSession, Task> OnGetUserSession;

    public UserRepo(IServiceProvider services) : base(services)
    {
    }

    protected override string GetHubName() => UserManagementHubMethodNames.HubName;

    protected override async Task HubConnectionOnReconnected(string? arg)
    {
        await Task.CompletedTask;
    }

    protected override async Task<IList<IDisposable>> On()
    {
        await Task.CompletedTask;

        IDisposable getGroups = HubConnection!.On<List<Group>>(UserManagementHubMethodNames.GetGroups, async (list) =>
        {
            await OnGetGroups?.Invoke(list);
        });

        IDisposable getGroupMemberships = HubConnection.On<List<GroupMembership>>(UserManagementHubMethodNames.GetGroupMemberships, async (list) =>
        {
            await OnGetGroupMemberships?.Invoke(list);
        });

        IDisposable getUsers = HubConnection.On<List<UserContainer>>(UserManagementHubMethodNames.GetUsers, async (list) =>
        {
            await OnGetUsers?.Invoke(list);
        });

        IDisposable getUserSession = HubConnection!.On<UserSession>(UserManagementHubMethodNames.GetUserSession, async (s) =>
        {
            await OnGetUserSession?.Invoke(s);
        });

        return new List<IDisposable>()
        {
            getGroups,
            getGroupMemberships,
            getUsers,
            getUserSession,
        };
    }

    public async Task NewGroup(Group group)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.NewGroup, group);
        }
    }

    public async Task UpdateGroup(Group group)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.UpdateGroup, group);
        }
    }

    public async Task DeleteGroup(Group group)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.DeleteGroup, group);
        }
    }

    public async Task AddUser(UserContainer user)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.AddUser, user);
        }
    }

    public async Task RemoveUser(UserContainer user)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.RemoveUser, user);
        }
    }

    public async Task UpdateUserSession(UserSession session)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(UserManagementHubMethodNames.UpdateUserSession, session);
            Console.WriteLine($"Getting session");
        }
    }
}
