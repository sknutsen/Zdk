using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Zdk.Server.Data;
using Zdk.Server.DataHandlers.Base;
using Zdk.Shared.Models;
using Zdk.Utilities.Authentication.Data;

namespace Zdk.Server.DataHandlers;

public class UserSessionHandler : HandlerBase<UserSessionHandler, DataContext>
{
    public UserSessionHandler(DataContext context, ILogger<UserSessionHandler> logger)
        : base(context, logger)
    {
    }

    public async Task<IList<UserSession>> List(int groupId)
    {
        var list = await List();

        return list.Where(e => e.GroupId == groupId).ToList();
    }

    public async Task<IList<UserSession>> List()
    {
        return await context.UserSessions.ToListAsync();
    }

    public async Task<UserSession> Get(string userId)
    {
        logger.LogInformation("Finding user session...");
        var session = await context.UserSessions.FindAsync(userId);

        if (session == null)
        {
            logger.LogInformation("No user session found. Creating new session...");
            session = new(userId);

            await context.UserSessions.AddAsync(session);
            await context.SaveChangesAsync();

            session = await context.UserSessions.FindAsync(userId);

            if (session == null)
            {
                throw new Exception();
            }
        }

        return session;
    }

    public async Task<bool> Update(UserSession session)
    {
        logger.LogInformation("Updating user session...");

        try
        {
            context.Update(session);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {
            return false;
        }
    }
}
