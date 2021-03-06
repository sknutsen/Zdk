namespace Zdk.Server.Hubs;

public partial class BaseHub<HubName>
{
    public async Task JoinGroup()
    {
        int groupId = await GetGroupId();

        await JoinGroup(groupId);
    }

    protected async Task JoinGroup(int groupId)
    {
        await Groups.AddToGroupAsync(Context.ConnectionId, groupId.ToString());
    }

    protected async Task LeaveGroup()
    {
        int groupId = await GetGroupId();

        await LeaveGroup(groupId);
    }

    public async Task LeaveGroup(int groupId)
    {
        await Groups.RemoveFromGroupAsync(Context.ConnectionId, groupId.ToString());
    }
}
