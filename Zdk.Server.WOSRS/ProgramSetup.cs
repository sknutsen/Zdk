using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Routing;
using Microsoft.Extensions.DependencyInjection;
using Zdk.Shared;

namespace Zdk.Server.WOSRS;

public static class ProgramSetup
{
    public static IMvcBuilder AddWOSRSApplicationPart(this IMvcBuilder builder)
    {
        return builder.AddApplicationPart(typeof(ProgramSetup).GetTypeInfo().Assembly);
    }
}
