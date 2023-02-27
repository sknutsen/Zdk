using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using OpenIddict.Validation.AspNetCore;
using System.Threading.Tasks;
using Zdk.DataAccess;
using Zdk.Utilities.Authentication.Data;
using Zdk.Utilities.Authentication.Helpers;

namespace Zdk.Server.WOSRS;

[Authorize(AuthenticationSchemes = OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme)]
[Route("[controller]")]
[ApiController]
public class UsersController : Controller
{
    private WOSRSContext context;
    private readonly ILogger<UsersController> logger;
    private readonly UserManager<ZdkUser> userManager;

    public UsersController(ILogger<UsersController> logger, UserManager<ZdkUser> userManager, WOSRSContext context)
    {
        this.logger = logger;
        this.userManager = userManager;
        this.context = context;
    }

    [HttpGet]
    public async Task<ActionResult<string>> Get()
    {
        var userId = User.GetUserId();

        return Ok(userId);
    }
}
