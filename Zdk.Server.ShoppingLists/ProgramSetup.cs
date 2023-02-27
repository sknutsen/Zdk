using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Zdk.Shared;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Routing;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared.Constants;

namespace Zdk.Server.ShoppingLists;

public static class ProgramSetup
{
    public static HubEndpointConventionBuilder MapShoppingListHubs(this IEndpointRouteBuilder builder)
    {
        return builder.MapHub<ShoppingListsHub>($"/{ShoppingListsHubMethodNames.HubName}");
    }
}
