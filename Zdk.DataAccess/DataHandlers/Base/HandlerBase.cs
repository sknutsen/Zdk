using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;

namespace Zdk.DataAccess;

public abstract class HandlerBase<Handler, Context> 
    where Context : DbContext
{
    protected readonly Context context;
    protected readonly ILogger<Handler> logger;

    public HandlerBase() { }
    
    public HandlerBase(Context context, ILogger<Handler> logger)
    {
        this.context = context;
        this.logger = logger;
    }
}
