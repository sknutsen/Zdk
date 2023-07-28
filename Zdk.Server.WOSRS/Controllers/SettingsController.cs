using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using OpenIddict.Validation.AspNetCore;
using System.Linq;
using System.Threading.Tasks;
using Zdk.DataAccess;
using Zdk.Shared.DataContainers;
using Zdk.Shared.Models;
using Zdk.Utilities.Authentication;

namespace Zdk.Server.WOSRS;

[Authorize(AuthenticationSchemes = OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme)]
[ApiController]
[Route("[controller]")]
public class SettingsController : Controller
{
    private WOSRSContext context;
    private readonly ILogger<SettingsController> logger;
    private readonly UserManager<ZdkUser> userManager;

    public SettingsController(ILogger<SettingsController> logger, UserManager<ZdkUser> userManager, WOSRSContext context)
    {
        this.logger = logger;
        this.userManager = userManager;
        this.context = context;
    }

    [HttpGet("get")]
    public async Task<ActionResult<SettingsContainer>> Get()
    {
        string userId = User.GetUserId();

        SettingsContainer result = context.Settings.Where(e => e.UserId == userId).Select(e => new SettingsContainer { SettingsId = e.SettingsId, OrderType = e.OrderType, PointSystem = e.PointSystem }).FirstOrDefault();

        if (result == null)
        {
            Settings newSettings = new Settings();
            newSettings.UserId = userId;
            newSettings.OrderType = 0;
            newSettings.PointSystem = false;

            var settings = await context.AddAsync(newSettings);
            await context.SaveChangesAsync();

            result = new SettingsContainer();
            result.Fill(settings.Entity);
        }

        return Ok(result);
    }

    [HttpPost("update")]
    public async Task<ActionResult<SettingsContainer>> Update([FromBody] SettingsContainer container)
    {
        string userId = User.GetUserId();

        Settings settings = await context.Settings.FindAsync(container.SettingsId);

        if (settings.UserId != userId)
        {
            return BadRequest();
        }

        settings.OrderType = container.OrderType;
        settings.PointSystem = container.PointSystem;

        await context.SaveChangesAsync();

        return NoContent();
    }
}
