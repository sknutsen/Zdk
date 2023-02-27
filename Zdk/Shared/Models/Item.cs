using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Zdk.Shared.Models;

[Table("items")]
public class Item : TimeStamped, IEntityClass
{
    [Key]
    [Column("item_id")]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ItemId { get; set; }

    [Required]
    [Column("item_name", TypeName = "text")]
    public string ItemName { get; set; }

    [Required]
    [ForeignKey("application_users")]
    [Column("user_id")]
    public string UserId { get; set; }

    public ICollection<ItemCategory> ItemCategories { get; set; }
    public ICollection<ScheduledItem> ScheduledItems { get; set; }
}
