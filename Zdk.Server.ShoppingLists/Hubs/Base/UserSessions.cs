using Zdk.Shared.Models;
using Zdk.Utilities.Authentication;

namespace Zdk.Server.ShoppingLists;

public partial class SlHubBase<HubName>
{
    protected override async Task<int> GetGroupId()
    {
        string userId = GetUserId();
        UserSession session = await userSessionHandler.Get(userId);

        return session.GroupId;
    }

    protected async Task<bool> IsAdmin()
    {
        string userId = GetUserId();

        return await userSessionHandler.IsAdmin(userId);
    }
}
