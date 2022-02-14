using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Models;

[Table("group_memberships")]
public class GroupMembership
{
    [Key]
    [Column("group_membership_id")]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int GroupMembershipId { get; set; }

    [Required]
    [ForeignKey("groups")]
    [Column("group_id")]
    public int GroupId { get; set; }

    [Required]
    [Column("user_id", TypeName = "text")]
    public string UserId { get; set; }

    [Required]
    [Column("is_current")]
    public bool IsCurrent { get; set; }
}
