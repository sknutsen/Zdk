using Microsoft.EntityFrameworkCore;
using Zdk.Shared.Models;

namespace Zdk.DataAccess;

public class WOSRSContext : DbContext
{
    public DbSet<Category> Categories { get; set; }
    public DbSet<Item> Items { get; set; }
    public DbSet<ItemCategory> ItemCategories { get; set; }
    public DbSet<ScheduledItem> ScheduledItems { get; set; }
    public DbSet<Settings> Settings { get; set; }

    public WOSRSContext(DbContextOptions<WOSRSContext> options) : base(options)
    {
    }

    protected override void OnModelCreating(ModelBuilder builder)
    {
        builder.Entity<ScheduledItem>()
            .HasOne<Item>(s => s.Item)
            .WithMany(i => i.ScheduledItems)
            .HasForeignKey(s => s.ItemId)
            .OnDelete(DeleteBehavior.Cascade);

        builder.Entity<ItemCategory>()
            .HasOne(ic => ic.Item)
            .WithMany(i => i.ItemCategories)
            .HasForeignKey(ic => ic.ItemId)
            .OnDelete(DeleteBehavior.Cascade);

        builder.Entity<ItemCategory>()
            .HasOne(ic => ic.Category)
            .WithMany(c => c.ItemCategories)
            .HasForeignKey(ic => ic.CategoryId)
            .OnDelete(DeleteBehavior.Cascade);

        base.OnModelCreating(builder);
    }
}
