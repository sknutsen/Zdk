using Microsoft.AspNetCore.Components;
using Zdk.Shared.Models;

namespace Zdk.Client
{
    public partial class UsersComponent : ZComponentBase
    {
        [Parameter]
        public IList<UserContainer>? Users { get; set; }

        [Parameter]
        public IList<GroupMembership>? GroupMemberships { get; set; }

        [Parameter]
        public int GroupId { get; set; }

        [Parameter]
        public Func<UserContainer, Task>? Add { get; set; }

        [Parameter]
        public Func<UserContainer, Task>? Remove { get; set; }

        private bool InGroup(string id)
        {
            return GroupMemberships.Any(e => e.UserId == id && e.GroupId == GroupId);
        }
    }
}