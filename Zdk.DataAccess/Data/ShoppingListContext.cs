using Microsoft.EntityFrameworkCore;
using Zdk.Shared.Models;

namespace Zdk.DataAccess;

public class ShoppingListContext : DbContext
{
    public DbSet<ShoppingListItem> Items { get; set; }
    public DbSet<Group> Groups { get; set; }
    public DbSet<ShoppingList> ShoppingLists { get; set; }
    public DbSet<GroupMembership> GroupMemberships { get; set; }
    public DbSet<UserSession> UserSessions { get; set; }

    public ShoppingListContext(DbContextOptions<ShoppingListContext> options)
        : base(options)
    {
    }

    protected override void OnModelCreating(ModelBuilder builder)
    {
        base.OnModelCreating(builder);

        //builder.Entity<ShoppingList>()
        //    .HasMany(e => e.Items)
        //    .WithOne()
        //    .HasForeignKey(e => e.ShoppingListId)
        //    .OnDelete(DeleteBehavior.Cascade);

        //builder.Entity<Group>()
        //    .HasMany(g => g.ShoppingLists)
        //    .WithOne()
        //    .HasForeignKey(e => e.GroupId)
        //    .OnDelete(DeleteBehavior.Cascade);

        //builder.Entity<GroupMembership>()
        //    .HasOne<Group>()
        //    .WithMany(e => e.GroupMemberships)
        //    .HasForeignKey(e => e.GroupId)
        //    .OnDelete(DeleteBehavior.Cascade);

        //builder.Entity<ShoppingList>()
        //    .HasOne<Group>()
        //    .WithMany(e => e.ShoppingLists)
        //    .HasForeignKey(e => e.GroupId)
        //    .OnDelete(DeleteBehavior.Cascade);

        //builder.Entity<Item>()
        //    .HasOne<ShoppingList>()
        //    .WithMany(e => e.Items)
        //    .HasForeignKey(e => e.ShoppingListId)
        //    .OnDelete(DeleteBehavior.Cascade);
    }
}
