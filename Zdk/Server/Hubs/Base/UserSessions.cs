using Zdk.Shared.Models;
using Zdk.Utilities.Authentication.Helpers;

namespace Zdk.Server.Hubs;

public partial class BaseHub<HubName>
{
    protected async Task<int> GetGroupId()
    {
        string userId = GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        return session.GroupId;
    }

    protected string GetUserId()
    {
        string userId = Context.User?.GetUserId() ?? "???";

        return userId;
    }

    protected async Task<bool> IsAdmin()
    {
        string userId = GetUserId();

        return await userSessionHandler.IsAdmin(userId);
    }
}
