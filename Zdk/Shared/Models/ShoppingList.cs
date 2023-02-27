using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Models;

[Table("shopping_lists")]
public class ShoppingList : EntityBase
{
    [Key]
    [Column("shopping_list_id")]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ShoppingListId { get; set; }

    [Required]
    [Column("shopping_list_name", TypeName = "text")]
    public string ShoppingListName { get; set; }

    [Required]
    [ForeignKey("groups")]
    [Column("group_id")]
    public int GroupId { get; set; }

    [Required]
    [Column("posted_by", TypeName = "text")]
    public string PostedBy { get; set; }

    public virtual ICollection<ShoppingListItem> Items { get; set; }
}
