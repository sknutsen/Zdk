using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using OpenIddict.Validation.AspNetCore;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Zdk.DataAccess;
using Zdk.Shared.Constants;
using Zdk.Shared.DataContainers;
using Zdk.Shared.Models;
using Zdk.Utilities.Authentication;

namespace Zdk.Server.WOSRS;

[Authorize(AuthenticationSchemes = OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme)]
[ApiController]
[Route("[controller]")]
public class ItemsController : Controller
{
    private WOSRSContext context;
    private readonly ILogger<ItemsController> logger;
    private readonly UserManager<ZdkUser> userManager;

    public ItemsController(ILogger<ItemsController> logger, UserManager<ZdkUser> userManager, WOSRSContext context)
    {
        this.logger = logger;
        this.userManager = userManager;
        this.context = context;
    }

    [HttpGet("list")]
    public async Task<ActionResult<IEnumerable<ItemContainer>>> List()
    {
        logger.LogInformation(LogTexts.ListingItems);

        var userId = User.GetUserId();

        var result = context.Items.Where(e => e.UserId == userId).Select<Item, ItemContainer>(e => new ItemContainer { ItemId = e.ItemId, ItemName = e.ItemName, ItemCategories = e.ItemCategories }).ToList();

        return Ok(result);
    }

    [HttpPost("get")]
    public async Task<ActionResult<ItemContainer>> Get([FromBody] ItemContainer container)
    {
        logger.LogInformation(LogTexts.GettingItem);

        var userId = User.GetUserId();

        var result = await context.Items.FindAsync(container.ItemId);

        if (result.UserId != userId)
        {
            return BadRequest();
        }

        return Ok(result);
    }

    [HttpPost("update")]
    public async Task<ActionResult<ItemContainer>> Update([FromBody] ItemContainer container)
    {
        logger.LogInformation(LogTexts.UpdatingItem);

        var userId = User.GetUserId();

        var item = await context.Items.FindAsync(container.ItemId);

        if (item.UserId != userId)
        {
            return BadRequest();
        }

        item.ItemName = container.ItemName;
        item.UpdatedAt = System.DateTime.Now;

        context.Update(item);
        await context.SaveChangesAsync();

        return NoContent();
    }

    [HttpPost("new")]
    public async Task<ActionResult<ItemContainer>> New([FromBody] ItemContainer container)
    {
        logger.LogInformation(LogTexts.CreatingItem);

        var userId = User.GetUserId();

        Item item = (Item)container.ToEntityClass();

        item.UserId = userId;
        item.CreatedAt = System.DateTime.Now;
        item.UpdatedAt = System.DateTime.Now;

        await context.AddAsync(item);
        await context.SaveChangesAsync();

        return NoContent();
    }

    [HttpPost("delete")]
    public async Task<ActionResult<ItemContainer>> Delete([FromBody] ItemContainer container)
    {
        logger.LogInformation(LogTexts.DeletingItem);

        var userId = User.GetUserId();

        var item = await context.Items.FindAsync(container.ItemId);

        if (item.UserId != userId)
        {
            return BadRequest();
        }

        context.Remove(item);
        await context.SaveChangesAsync();

        return NoContent();
    }
}
