using Microsoft.EntityFrameworkCore;
using Zdk.Server.Data;
using Zdk.Server.DataHandlers.Base;
using Zdk.Shared.Models;

namespace Zdk.Server.DataHandlers;

public class GroupHandler : HandlerBase<GroupHandler, DataContext>
{
    public GroupHandler(DataContext context, ILogger<GroupHandler> logger) 
        : base(context, logger)
    {
    }

    public async Task<IList<Group>> List(string userId, bool fullList = false)
    {
        IList<Group> list = await List();

        if (fullList)
        {
            return list;
        }
        else
        {
            IList<GroupMembership> memberships = await context.GroupMemberships.Where(e => e.UserId == userId).ToListAsync() ?? new List<GroupMembership>();

            return list.Where(e => memberships.Any(f => f.GroupId == e.GroupId)).ToList() ?? new List<Group>();
        }
    }

    public async Task<IList<Group>> List()
    {
        return await context.Groups.Include(e => e.GroupMemberships).Include(e => e.ShoppingLists).ToListAsync() ?? new List<Group>();
    }

    public async Task<Group?> Get(int id)
    {
        return await context.Groups.FindAsync(id);
    }

    public async Task<Group> Get(string name, string owner)
    {
        return await context.Groups.Where(e => e.GroupName == name && e.Owner == owner).FirstAsync();
    }

    public async Task<(bool, Group?)> New(Group group)
    {
        try
        {
            bool exists = context.Groups.Any(e => e.GroupName == group.GroupName && e.Owner == group.Owner);

            if (exists)
            {
                throw new Exception("Group exists");
            }

            group.CreatedAt = DateTime.Now;
            group.UpdatedAt = DateTime.Now;
            var result = await context.Groups.AddAsync(group);
            await context.SaveChangesAsync();

            return (true, result.Entity);
        }
        catch (Exception e)
        {
            logger.LogError(e.Message);

            return (false, null);
        }
    }

    public async Task<bool> Update(Group group)
    {
        try
        {
            group.UpdatedAt = DateTime.Now;
            context.Groups.Update(group);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Delete(Group group)
    {
        try
        {
            context.Groups.Remove(group);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }
}
