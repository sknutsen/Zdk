using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace Zdk.DataAccess;

public static class ProgramSetup
{
    public static IServiceCollection AddZdkDbContexts(this IServiceCollection services, string shoppingListConnectionString, string wosrsConnectionString)
    {
        return services.AddDbContext<ShoppingListContext>(options =>
                       {
                           options.UseNpgsql(shoppingListConnectionString, o => o.SetPostgresVersion(new Version(9, 6)));
                       })
                       .AddDbContext<WOSRSContext>(options =>
                       {
                           options.UseNpgsql(wosrsConnectionString, o => o.SetPostgresVersion(new Version(9, 6)));
                       });
    }
}
