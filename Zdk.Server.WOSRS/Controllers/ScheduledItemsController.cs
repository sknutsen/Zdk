using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using OpenIddict.Validation.AspNetCore;
using System;
using System.Collections;
using System.Collections.Generic;
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
public class ScheduledItemsController : Controller
{
    private WOSRSContext context;
    private readonly ILogger<ScheduledItemsController> logger;
    private readonly UserManager<ZdkUser> userManager;

    public ScheduledItemsController(ILogger<ScheduledItemsController> logger, UserManager<ZdkUser> userManager, WOSRSContext context)
    {
        this.logger = logger;
        this.userManager = userManager;
        this.context = context;
    }

    [HttpGet("current")]
    public async Task<ActionResult<ScheduledItemContainer>> Current()
    {
        logger.LogInformation("Getting current item...");

        var userId = User.GetUserId();
        var today = DateTime.Today;

        var result = context.ScheduledItems.Where(e => e.UserId == userId && e.Date.Day == today.Day && e.Date.Month == today.Month && e.Date.Year == today.Year).FirstOrDefault();

        if (result == null)
        {
            return NoContent();
        }

        return Ok(new ScheduledItemContainer(result));
    }

    [HttpGet("list")]
    public async Task<ActionResult<IEnumerable<ScheduledItemContainer>>> List()
    {
        logger.LogInformation("Listing items...");

        var userId = User.GetUserId();

        var result = context.ScheduledItems.Where(e => e.UserId == userId)
                                        .Select<ScheduledItem, ScheduledItemContainer>(e => new ScheduledItemContainer
                                        {
                                            ScheduledItemId = e.ScheduledItemId,
                                            Date = e.Date,
                                            ScheduleGroup = e.ScheduleGroup,
                                            IsComplete = e.IsComplete,
                                            ItemId = e.ItemId,
                                            Item = e.Item
                                        })
                                        .OrderByDescending(e => e.Date)
                                        .ToList();

        return Ok(result);
    }

    [HttpPost("get")]
    public async Task<ActionResult<ScheduledItemContainer>> Get([FromBody] ScheduledItemContainer container)
    {
        logger.LogInformation("Getting item...");

        var userId = User.GetUserId();

        var result = await context.ScheduledItems.FindAsync(container.ScheduledItemId);

        if (result.UserId != userId)
        {
            return BadRequest();
        }

        return Ok(new ScheduledItemContainer(result));
    }

    [HttpPost("update")]
    public async Task<ActionResult<ScheduledItemContainer>> Update([FromBody] ScheduledItemContainer container)
    {
        logger.LogInformation("Updating item...");

        var userId = User.GetUserId();

        var item = await context.ScheduledItems.FindAsync(container.ScheduledItemId);

        if (item.UserId != userId)
        {
            return BadRequest();
        }

        item.IsComplete = container.IsComplete;
        item.UpdatedAt = DateTime.Now;

        context.Update(item);
        await context.SaveChangesAsync();

        return NoContent();
    }

    [HttpPost("new")]
    public async Task<ActionResult<ScheduledItemContainer>> New([FromBody] ScheduledItemContainer container)
    {
        logger.LogInformation("Creating item...");

        var userId = User.GetUserId();

        ScheduledItem scheduledItem = (ScheduledItem)container.ToEntityClass();

        var item = await context.FindAsync<Item>(container.ItemId);

        scheduledItem.Item = item;

        scheduledItem.UserId = userId;
        scheduledItem.CreatedAt = DateTime.Now;
        scheduledItem.UpdatedAt = DateTime.Now;
        scheduledItem.Date = DateTime.Now;

        await context.AddAsync(scheduledItem);
        await context.SaveChangesAsync();

        return NoContent();
    }

    [HttpPost("delete")]
    public async Task<ActionResult<ScheduledItemContainer>> Delete([FromBody] ScheduledItemContainer container)
    {
        logger.LogInformation("Deleting item...");

        var userId = User.GetUserId();

        var item = await context.ScheduledItems.FindAsync(container.ScheduledItemId);

        if (item.UserId != userId)
        {
            return BadRequest();
        }

        context.Remove(item);
        await context.SaveChangesAsync();

        return NoContent();
    }
}
