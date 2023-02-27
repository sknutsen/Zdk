using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using Zdk.Shared.Models;

namespace Zdk.DataAccess;

public class GroupMembershipHandler : HandlerBase<GroupMembershipHandler, ShoppingListContext>
{
    public GroupMembershipHandler(ShoppingListContext context, ILogger<GroupMembershipHandler> logger) 
        : base(context, logger)
    {
    }

    public async Task<IList<GroupMembership>> List(int groupId)
    {
        return await context.GroupMemberships.Where(e => e.GroupId == groupId).ToListAsync();
    }

    public async Task<IList<GroupMembership>> List(string userId)
    {
        return await context.GroupMemberships.Where(e => e.UserId == userId).ToListAsync();
    }

    public async Task<GroupMembership?> Get(int id)
    {
        return await context.GroupMemberships.FindAsync(id);
    }

    public async Task<(bool, GroupMembership?)> New(int groupId, string userId, bool isCurrent)
    {
        GroupMembership membership = new()
        {
            GroupId = groupId,
            UserId = userId,
            IsCurrent = isCurrent,
        };

        return await New(membership);
    }

    public async Task<(bool, GroupMembership?)> New(GroupMembership membership)
    {
        try
        {
            if (Exists(membership))
            {
                throw new Exception("User is already member of group");
            }

            var newMembership = await context.GroupMemberships.AddAsync(membership);
            await context.SaveChangesAsync();

            return (true, newMembership.Entity);
        }
        catch (Exception)
        {

            return (false, null);
        }
    }

    public async Task<bool> Delete(int groupMembershipId)
    {
        try
        {
            var membership = await context.GroupMemberships.FindAsync(groupMembershipId);

            if (membership is null)
            {
                throw new Exception($"Invalid id for GroupMembership. Id: {groupMembershipId}");
            }

            return await Delete(membership);
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Delete(int groupId, string userId)
    {
        try
        {
            var membership = context.GroupMemberships.Where(e => e.GroupId == groupId && e.UserId == userId).FirstOrDefault();

            if (membership is null)
            {
                throw new Exception($"Invalid group and user combo");
            }

            return await Delete(membership);
        }
        catch (Exception)
        {

            return false;
        }
    }

    public async Task<bool> Delete(GroupMembership membership)
    {
        try
        {
            context.GroupMemberships.Remove(membership);
            await context.SaveChangesAsync();

            return true;
        }
        catch (Exception)
        {

            return false;
        }
    }

    public bool Exists(int groupId, string userId)
    {
        GroupMembership membership = new GroupMembership
        {
            GroupId = groupId,
            UserId = userId,
        };
        return Exists(membership);
    }

    public bool Exists(GroupMembership membership)
    {
        return context.GroupMemberships.Any(e => e.UserId == membership.UserId && e.GroupId == membership.GroupId);
    }
}
