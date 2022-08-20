using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Models;

[Table("groups")]
public class Group : EntityBase
{
    [Key]
    [Column("group_id")]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int GroupId { get; set; }

    [Required]
    [Column("group_name", TypeName = "text")]
    public string GroupName { get; set; }

    [Required]
    [Column("owner", TypeName = "text")]
    public string Owner { get; set; }

    public virtual ICollection<GroupMembership> GroupMemberships { get; set; }
    public virtual ICollection<ShoppingList> ShoppingLists { get; set; }
}
