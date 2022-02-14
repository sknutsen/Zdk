using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Zdk.Shared.Models.Base;

namespace Zdk.Shared.Models;

[Table("items")]
public class Item : EntityBase
{
    [Key]
    [Column("item_id")]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ItemId { get; set; }

    [Required]
    [Column("item_name", TypeName = "text")]
    public string ItemName { get; set; }

    [Required]
    [Column("amount")]
    public int Amount { get; set; }

    [Required]
    [Column("sold_out")]
    public bool SoldOut { get; set; }

    [Required]
    [ForeignKey("shopping_lists")]
    [Column("shopping_list_id")]
    public int ShoppingListId { get; set; }

    [Required]
    [Column("posted_by", TypeName = "text")]
    public string PostedBy { get; set; }
}
