using Microsoft.AspNetCore.SignalR;
using Zdk.Server.DataHandlers;

namespace Zdk.Server.Hubs;

public abstract partial class BaseHub<HubName> : Hub
{
    protected readonly GroupHandler groupHandler;
    protected readonly GroupMembershipHandler groupMembershipHandler;
    protected readonly UserSessionHandler userSessionHandler;
    protected readonly ILogger<HubName> logger;

    public BaseHub(GroupHandler groupHandler, GroupMembershipHandler groupMembershipHandler, UserSessionHandler userSessionHandler, ILogger<HubName> logger)
    {
        this.groupHandler = groupHandler;
        this.groupMembershipHandler = groupMembershipHandler;
        this.userSessionHandler = userSessionHandler;
        this.logger = logger;
    }
}
