using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Zdk.DataAccess;
using Zdk.Shared.Models;
using Zdk.Utilities.SignalR;

namespace Zdk.Server.ShoppingLists;

public partial class SlHubBase<HubName> : BaseHub<HubName>
{
    protected readonly GroupHandler groupHandler;
    protected readonly GroupMembershipHandler groupMembershipHandler;
    protected readonly UserSessionHandler userSessionHandler;

    public SlHubBase(IServiceProvider services) : base(services)
    {
        groupHandler = services.GetRequiredService<GroupHandler>();
        groupMembershipHandler = services.GetRequiredService<GroupMembershipHandler>();
        userSessionHandler = services.GetRequiredService<UserSessionHandler>();
    }
}
