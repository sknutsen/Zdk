using Azure.Identity;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.ResponseCompression;
using Microsoft.AspNetCore.SignalR;
using Azure.Security.KeyVault;
using Microsoft.EntityFrameworkCore;
using System.Security.Cryptography.X509Certificates;
using Zdk.Utilities.Authentication.Data;
using Zdk.Server;
using Zdk.DataAccess;
using Zdk.Server.Hubs;
using Zdk.Server.ShoppingLists;
using static OpenIddict.Abstractions.OpenIddictConstants;
using Azure.Security.KeyVault.Secrets;
using Azure.Security.KeyVault.Certificates;
using Zdk.Server.WOSRS;

WebApplicationBuilder builder = WebApplication.CreateBuilder(args);

builder.Host.ConfigureLogging(logging =>
{
    //logging.ClearProviders();
    //logging.AddEventLog();
    logging.AddConsole();
    logging.AddDebug();
    logging.AddAzureWebAppDiagnostics();
    logging.AddEventSourceLogger();
});

string vaultUri = builder.Configuration["VaultUri"];

bool isDev = builder.Environment.IsDevelopment();

builder.Host.ConfigureAppConfiguration((context, config) =>
{
    if (!isDev)
    {
        Uri keyVaultEndpoint = new(vaultUri);
        config.AddAzureKeyVault(keyVaultEndpoint, new DefaultAzureCredential());
    }
});
DefaultAzureCredential cred = new DefaultAzureCredential();
SecretClient client = new(new Uri(vaultUri), cred);
CertificateClient certClient = new(new Uri(vaultUri), cred);
string certName = builder.Configuration["Cert"];
KeyVaultSecret certificate = await client.GetSecretAsync(certName);

byte[] privateKeyBytes = Convert.FromBase64String(certificate.Value);
X509Certificate2 certificateWithPrivateKey = new(privateKeyBytes, (string)null, X509KeyStorageFlags.MachineKeySet);

string zdkAuthConnectionString = builder.Configuration["ConnectionStrings:ZdkAuthDb"];
string shoppingListConnectionString = builder.Configuration["ConnectionStrings:DataDb"];
string wosrsConnectionString = builder.Configuration["ConnectionStrings:wosrsDb"];

if (string.IsNullOrEmpty(zdkAuthConnectionString))
{
    zdkAuthConnectionString = builder.Configuration["ZdkAuthDb"];
}

if (string.IsNullOrEmpty(shoppingListConnectionString))
{
    shoppingListConnectionString = builder.Configuration["DataDb"];
}

if (string.IsNullOrEmpty(wosrsConnectionString))
{
    wosrsConnectionString = builder.Configuration["wosrsDb"];
}

builder.Services.AddZdkDbContexts(shoppingListConnectionString, wosrsConnectionString);

builder.Services.AddDbContext<AuthContext>(options => 
{
    options.UseNpgsql(zdkAuthConnectionString, o => o.SetPostgresVersion(new Version(9, 6)));
    options.UseOpenIddict();
});

//builder.Services.AddDbContext<ShoppingListContext>(options => 
//{
//    options.UseNpgsql(shoppingListConnectionString, o => o.SetPostgresVersion(new Version(9, 6)));
//});

builder.Services.AddDefaultIdentity<ZdkUser>(options => options.SignIn.RequireConfirmedAccount = false)
                .AddEntityFrameworkStores<AuthContext>()
                .AddDefaultTokenProviders();

builder.Services.AddOpenIddict()
                .AddCore(options =>
                {
                    options.UseEntityFrameworkCore().UseDbContext<AuthContext>();
                })
                .AddServer(options =>
                {
                    // Enable the authorization, logout, token and userinfo endpoints.
                    options.SetAuthorizationEndpointUris("/connect/authorize")
                           .SetLogoutEndpointUris("/connect/logout")
                           .SetTokenEndpointUris("/connect/token")
                           .SetUserinfoEndpointUris("/connect/userinfo");

                    // Mark the "email", "profile" and "roles" scopes as supported scopes.
                    options.RegisterScopes(Scopes.Email, Scopes.Profile, Scopes.Roles);

                    // Note: the sample uses the code and refresh token flows but you can enable
                    // the other flows if you need to support implicit, password or client credentials.
                    options.AllowAuthorizationCodeFlow()
                           .AllowRefreshTokenFlow();

                    // Register the signing and encryption credentials.
                    options.AddEncryptionCertificate(certificateWithPrivateKey)
                           .AddSigningCertificate(certificateWithPrivateKey);

                    // Register the ASP.NET Core host and configure the ASP.NET Core-specific options.
                    options.UseAspNetCore()
                           .EnableAuthorizationEndpointPassthrough()
                           .EnableLogoutEndpointPassthrough()
                           .EnableStatusCodePagesIntegration()
                           .EnableTokenEndpointPassthrough();
                })
                .AddValidation(options =>
                {
                    options.UseLocalServer();

                    options.UseAspNetCore();
                });

builder.Services.AddAuthentication();
builder.Services.AddAuthorization();

builder.Services.AddSignalR();

builder.Services.AddControllers()
                .AddWOSRSApplicationPart()
                .AddControllersAsServices();

builder.Services.AddControllersWithViews();
builder.Services.AddRazorPages();
builder.Services.AddResponseCompression(opts =>
{
    opts.MimeTypes = ResponseCompressionDefaults.MimeTypes.Concat(new[] { "application/octet-stream" });
});

builder.Services.AddSingleton<IUserIdProvider, NameUserIdProvider>();

builder.Services.Configure<IdentityOptions>(options =>
{
    options.ClaimsIdentity.UserNameClaimType = Claims.Name;
    options.ClaimsIdentity.UserIdClaimType = Claims.Subject;
    options.ClaimsIdentity.RoleClaimType = Claims.Role;

    // Password settings.
    options.Password.RequireDigit = false;
    options.Password.RequireLowercase = false;
    options.Password.RequireNonAlphanumeric = false;
    options.Password.RequireUppercase = false;
    options.Password.RequiredLength = 6;
    options.Password.RequiredUniqueChars = 1;

    // User settings.
    options.User.AllowedUserNameCharacters =
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._@+";
    options.User.RequireUniqueEmail = false;
});

builder.Services.AddHostedService<ZdkWorker>();

builder.Services.AddScoped<ItemHandler>();
builder.Services.AddScoped<ShoppingListHandler>();
builder.Services.AddScoped<GroupHandler>();
builder.Services.AddScoped<GroupMembershipHandler>();
builder.Services.AddScoped<UserHandler>();
builder.Services.AddScoped<UserSessionHandler>();

WebApplication app = builder.Build();

app.UseResponseCompression();

// Configure the HTTP request pipeline.
if (isDev)
{
    app.UseDeveloperExceptionPage();
    app.UseMigrationsEndPoint();
    app.UseWebAssemblyDebugging();
}
else
{
    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseBlazorFrameworkFiles();
app.UseStaticFiles();

app.UseRouting();

app.UseAuthentication();
app.UseAuthorization();

app.UseEndpoints(endpoints =>
{
    endpoints.MapRazorPages();
    endpoints.MapControllers();
    endpoints.MapHub<MainHub>("/mainhub");
    endpoints.MapShoppingListHubs();
    endpoints.MapHub<UserManagementHub>("/userhub");
    endpoints.MapFallbackToFile("index.html");
});

app.Run();
