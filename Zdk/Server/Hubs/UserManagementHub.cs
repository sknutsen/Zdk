using Microsoft.AspNetCore.Authorization;
using Zdk.DataAccess;
using Zdk.Utilities.SignalR;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    private readonly GroupHandler groupHandler;
    private readonly GroupMembershipHandler groupMembershipHandler;
    private readonly UserHandler userHandler;
    private readonly UserSessionHandler userSessionHandler;

    public UserManagementHub(IServiceProvider services) : base(services)
    {
        groupHandler = services.GetRequiredService<GroupHandler>();
        groupMembershipHandler = services.GetRequiredService<GroupMembershipHandler>();
        userHandler = services.GetRequiredService<UserHandler>();
        userSessionHandler = services.GetRequiredService<UserSessionHandler>();
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
