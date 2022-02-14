using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using Zdk.Utilities.Authentication.Data;

namespace Zdk.Server.Data;

public class ZdkAuthContext : IdentityDbContext<ZdkUser>
{
    public ZdkAuthContext(DbContextOptions<ZdkAuthContext> options)
        : base(options)
    {
    }

    protected override void OnModelCreating(ModelBuilder builder)
    {
        base.OnModelCreating(builder);
    }
}
