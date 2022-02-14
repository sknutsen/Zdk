using Microsoft.EntityFrameworkCore;
using Zdk.Shared.Models;

namespace Zdk.Server.Data;

public class DataContext : DbContext
{
    public DbSet<Item> Items { get; set; }
    public DbSet<Group> Groups { get; set; }
    public DbSet<ShoppingList> ShoppingLists { get; set; }
    public DbSet<GroupMembership> GroupMemberships { get; set; }
    public DbSet<UserSession> UserSessions { get; set; }
 
    public DataContext(DbContextOptions<DataContext> options) 
        : base(options)
    {
    }

    protected override void OnModelCreating(ModelBuilder builder)
    {
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
        
        base.OnModelCreating(builder);
    }
}
