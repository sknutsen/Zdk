using OpenIddict.Abstractions;
using Zdk.Utilities.Authentication.Data;
using Zdk.Utilities.Authentication;
using static OpenIddict.Abstractions.OpenIddictConstants;

namespace Zdk.Server;

public class ZdkWorker : Worker
{
    private readonly IServiceProvider _serviceProvider;

    public ZdkWorker(IServiceProvider serviceProvider) : base(serviceProvider)
        => _serviceProvider = serviceProvider;

    protected override void SetDescriptor()
    {
        openIddictApplicationDescriptor = new OpenIddictApplicationDescriptor
        {
            ClientId = clientId,
            ConsentType = ConsentTypes.Explicit,
            DisplayName = displayName,
            Type = ClientTypes.Public,
            PostLogoutRedirectUris =
            {
                new Uri("https://www.zdk.no/authentication/logout-callback"),
                new Uri("https://zdk.no/authentication/logout-callback"),
                new Uri("https://localhost:44367/authentication/logout-callback"),
                new Uri("https://localhost:7203/authentication/logout-callback"),
                new Uri("https://zdkserver.azurewebsites.net/authentication/logout-callback"),
            },
            RedirectUris =
            {
                new Uri("https://www.zdk.no/authentication/login-callback"),
                new Uri("https://zdk.no/authentication/login-callback"),
                new Uri("https://localhost:44367/authentication/login-callback"),
                new Uri("https://localhost:7203/authentication/login-callback"),
                new Uri("https://zdkserver.azurewebsites.net/authentication/login-callback"),
            },
            Permissions =
            {
                Permissions.Endpoints.Authorization,
                Permissions.Endpoints.Logout,
                Permissions.Endpoints.Token,
                Permissions.GrantTypes.AuthorizationCode,
                Permissions.GrantTypes.RefreshToken,
                Permissions.ResponseTypes.Code,
                Permissions.Scopes.Email,
                Permissions.Scopes.Profile,
                Permissions.Scopes.Roles
            },
            Requirements =
            {
                Requirements.Features.ProofKeyForCodeExchange
            }
        };
    }

    protected override void SetValues()
    {
        clientId = "Zdk.Client";
        displayName = "Zdk";
    }
}
