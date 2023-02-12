using Microsoft.AspNetCore.Authorization;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

[Authorize]
public partial class UserManagementHub : BaseHub<UserManagementHub>
{
    private readonly UserHandler userHandler;

    public UserManagementHub(IServiceProvider services) : base(services)
    {
        userHandler = services.GetRequiredService<UserHandler>();
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
