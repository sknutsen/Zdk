using System;
using Zdk.Shared.Models;

namespace Zdk.Shared.DataContainers;

public class ScheduledItemContainer : ContainerBase, IContainer
{
    public int ScheduledItemId { get; set; }
    public bool IsComplete { get; set; }
    public DateTime Date { get; set; }
    public int? ScheduleGroup { get; set; }
    public int ItemId { get; set; }

    public Item Item { get; set; }

    public ScheduledItemContainer() { }

    public ScheduledItemContainer(IEntityClass entity)
    {
        Fill(entity);
    }

    public IEntityClass ToEntityClass()
    {
        return new ScheduledItem
        {
            ScheduledItemId = ScheduledItemId,
            IsComplete = IsComplete,
            Date = Date,
            ScheduleGroup = ScheduleGroup,
            ItemId = ItemId,
            Item = Item
        };
    }

    public void Fill(IEntityClass entity)
    {
        ScheduledItem scheduledItem = (ScheduledItem)entity;

        ScheduledItemId = scheduledItem.ScheduledItemId;
        IsComplete = scheduledItem.IsComplete;
        Date = scheduledItem.Date;
        ScheduleGroup = scheduledItem.ScheduleGroup;
        ItemId = scheduledItem.ItemId;
        Item = scheduledItem.Item;
    }
}
