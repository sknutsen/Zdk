using Microsoft.EntityFrameworkCore;
using Zdk.Server.Data;

namespace Zdk.Server.DataHandlers.Base;

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
