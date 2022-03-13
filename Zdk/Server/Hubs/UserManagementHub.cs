using Microsoft.AspNetCore.Authorization;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    private readonly UserHandler userHandler;

    public UserManagementHub(GroupHandler groupHandler, GroupMembershipHandler groupMembershipHandler, UserHandler userHandler, UserSessionHandler userSessionHandler, ILogger<UserManagementHub> logger)
        : base(groupHandler, groupMembershipHandler, userSessionHandler, logger)
    {
        this.userHandler = userHandler;
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
}
