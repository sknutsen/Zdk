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
using Zdk.Utilities.Authentication.Data;
using Zdk.Utilities.Authentication.Helpers;

namespace Zdk.Server.WOSRS;

[Authorize(AuthenticationSchemes = OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme)]
[ApiController]
[Route("[controller]")]
public class CategoriesController : Controller
{
    private WOSRSContext context;
    private readonly ILogger<CategoriesController> logger;
    private readonly UserManager<ZdkUser> userManager;

    public CategoriesController(ILogger<CategoriesController> logger, UserManager<ZdkUser> userManager, WOSRSContext context)
    {
        this.logger = logger;
        this.userManager = userManager;
        this.context = context;
    }

    [HttpGet("list")]
    public async Task<ActionResult<IEnumerable<CategoryContainer>>> List()
    {
        logger.LogInformation(LogTexts.ListingCategories);

        var userId = User.GetUserId();

        var list = context.Categories.Where(e => e.UserId == userId).Select<Category, CategoryContainer>(e => new CategoryContainer { CategoryId = e.CategoryId, CategoryName = e.CategoryName, ItemCategories = e.ItemCategories }).ToList();

        logger.LogInformation(LogTexts.ListCategoriesSuccess);

        return Ok(list);
    }

    [HttpPost("get")]
    public async Task<ActionResult<CategoryContainer>> Get([FromBody] CategoryContainer container)
    {
        logger.LogInformation(LogTexts.GettingCategory);

        var userId = User.GetUserId();

        var result = await context.Categories.FindAsync(container.CategoryId);

        if (result.UserId != userId)
        {
            logger.LogInformation(LogTexts.GetCategoryFailed);

            return BadRequest();
        }

        logger.LogInformation(LogTexts.GetCategorySuccess);

        return Ok(result);
    }

    [HttpPost("update")]
    public async Task<ActionResult<CategoryContainer>> Update([FromBody] CategoryContainer container)
    {
        logger.LogInformation(LogTexts.UpdatingCategory);

        var userId = User.GetUserId();

        var item = await context.Categories.FindAsync(container.CategoryId);

        if (item.UserId != userId)
        {
            logger.LogInformation(LogTexts.UpdateCategoryFailed);

            return BadRequest();
        }

        item.CategoryName = container.CategoryName;
        item.UpdatedAt = System.DateTime.Now;

        context.Update(item);
        await context.SaveChangesAsync();

        logger.LogInformation(LogTexts.UpdateCategorySuccess);

        return NoContent();
    }

    [HttpPost("new")]
    public async Task<ActionResult<CategoryContainer>> New([FromBody] CategoryContainer container)
    {
        logger.LogInformation(LogTexts.CreatingCategory);

        var userId = User.GetUserId();

        Category category = (Category)container.ToEntityClass();

        category.UserId = userId;
        category.CreatedAt = System.DateTime.Now;
        category.UpdatedAt = System.DateTime.Now;

        await context.AddAsync(category);
        await context.SaveChangesAsync();

        logger.LogInformation(LogTexts.NewCategorySuccess);

        return NoContent();
    }

    [HttpPost("delete")]
    public async Task<ActionResult<CategoryContainer>> Delete([FromBody] CategoryContainer container)
    {
        logger.LogInformation(LogTexts.DeletingCategory);

        var userId = User.GetUserId();

        var category = await context.Categories.FindAsync(container.CategoryId);

        if (category.UserId != userId)
        {
            logger.LogInformation(LogTexts.DeleteCategoryFailed);

            return BadRequest();
        }

        context.Remove(category);
        await context.SaveChangesAsync();

        logger.LogInformation(LogTexts.DeleteCategorySuccess);

        return NoContent();
    }
}
