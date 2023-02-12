using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Zdk.Client;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.RootComponents.Add<App>("#app");
//builder.RootComponents.Add<HeadOutlet>("head::after");

var baseUri = builder.HostEnvironment.BaseAddress;

builder.Services.AddHttpClient("Zdk.ServerAPI")
                .ConfigureHttpClient(client => client.BaseAddress = new Uri(baseUri))
                .AddHttpMessageHandler<BaseAddressAuthorizationMessageHandler>();

// Supply HttpClient instances that include access tokens when making requests to the server project
builder.Services.AddScoped(provider =>
{
    var factory = provider.GetRequiredService<IHttpClientFactory>();
    return factory.CreateClient("Zdk.ServerAPI");
});

builder.Services.AddScoped<ShoppingListsRepo>();
builder.Services.AddScoped<UserRepo>();

builder.Services.AddOptions();
builder.Services.AddOidcAuthentication(options =>
{
    options.ProviderOptions.ClientId = "Zdk.Client";
    options.ProviderOptions.Authority = baseUri;
    options.ProviderOptions.RedirectUri = $"{baseUri}authentication/login-callback";
    options.ProviderOptions.ResponseMode = "query";
    options.ProviderOptions.ResponseType = "code";

    options.AuthenticationPaths.LogInCallbackPath = $"{baseUri}authentication/login-callback";
    options.AuthenticationPaths.LogOutCallbackPath = $"{baseUri}authentication/logout-callback";
    // options.AuthenticationPaths.RemoteRegisterPath = $"{baseUri}Identity/Account/Register";
});

await builder.Build().RunAsync();
