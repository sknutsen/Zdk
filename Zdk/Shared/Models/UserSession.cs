using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Models;

[Table("user_sessions")]
public class UserSession
{
    [Key]
    [Column("user_id")]
    public string UserId { get; set; }

    [Required]
    [Column("group_id")]
    public int GroupId { get; set; }

    public UserSession() { }

    public UserSession(string userId)
    {
        UserId = userId;
        GroupId = 0;
    }

    public UserSession(string userId, int groupId)
    {
        UserId = userId;
        GroupId = groupId;
    }
}
