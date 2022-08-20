using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using System.Security.Claims;
using Zdk.Server.Data;
using Zdk.Utilities.Authentication.Data;

namespace Zdk.Server.DataHandlers;

public class UserHandler : HandlerBase<UserHandler, AuthContext>
{
    private UserManager<ZdkUser> userManager;

    public UserHandler(AuthContext context, ILogger<UserHandler> logger, UserManager<ZdkUser> userManager)
        : base(context, logger)
    {
        this.userManager = userManager;
    }

    public async Task<IList<ZdkUser>> List()
    {
        var users = await context.Users.ToListAsync();

        return users;
    }

    public async Task<IList<ZdkUser>> List(string role)
    {
        var users = await userManager.GetUsersInRoleAsync(role);

        return users;
    }

    public async Task<ZdkUser?> Get(string userId)
    {
        return await userManager.FindByIdAsync(userId);
    }

    public async Task<bool> UserHasRole(string userId, string role)
    {
        bool result = false;

        ZdkUser? user = await Get(userId);

        if (user is null)
        {
            return result;
        }

        result = await userManager.IsInRoleAsync(user, role);

        return result;
    }
}
