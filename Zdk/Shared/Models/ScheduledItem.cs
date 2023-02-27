using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Models;

[Table("scheduled_items")]
public class ScheduledItem : TimeStamped, IEntityClass
{
    [Key]
    [Column("scheduled_item_id")]
    public int ScheduledItemId { get; set; }

    [Column("is_complete")]
    public bool IsComplete { get; set; }

    [Column("date", TypeName = "date")]
    public DateTime Date { get; set; }

    [Column("schedule_group")]
    public int? ScheduleGroup { get; set; }

    [Required]
    [ForeignKey("items")]
    [Column("item_id")]
    public int ItemId { get; set; }
    public Item Item { get; set; }

    [Required]
    [ForeignKey("application_users")]
    [Column("user_id")]
    public string UserId { get; set; }
}
