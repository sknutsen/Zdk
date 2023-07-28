using System.Collections.Generic;
using OpenIddict.Abstractions;
using Zdk.Utilities.Authentication;
using static OpenIddict.Abstractions.OpenIddictConstants;

namespace Zdk.Server;

public class ZWorker : ZdkWorker
{
    public ZWorker(IServiceProvider serviceProvider) 
        : base(serviceProvider, 
               "Zdk.Client", 
               "Zdk", 
               new HashSet<Uri>{
                   new Uri("https://www.zdk.no/authentication/logout-callback"),
                   new Uri("https://zdk.no/authentication/logout-callback"),
                   new Uri("https://localhost:44367/authentication/logout-callback"),
                   new Uri("https://localhost:7203/authentication/logout-callback"),
                   new Uri("https://zdkserver.azurewebsites.net/authentication/logout-callback"),
               }, 
               new HashSet<Uri>
               {
                   new Uri("https://www.zdk.no/authentication/login-callback"),
                   new Uri("https://zdk.no/authentication/login-callback"),
                   new Uri("https://localhost:44367/authentication/login-callback"),
                   new Uri("https://localhost:7203/authentication/login-callback"),
                   new Uri("https://zdkserver.azurewebsites.net/authentication/login-callback"),
               }, 
               new HashSet<string>
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
               new HashSet<string>
               {
                   Requirements.Features.ProofKeyForCodeExchange
               })
    {
    }
}
