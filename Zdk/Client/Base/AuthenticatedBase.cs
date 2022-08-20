using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Authorization;

namespace Zdk.Client;

public abstract class AuthenticatedBase : ZComponentBase
{
    [CascadingParameter]
    public AuthenticationState AuthenticationState { get; set; }
    
    [Inject]
    protected AuthenticationStateProvider AuthenticationStateProvider { get; set; }
    
    [Inject]
    protected NavigationManager NavigationManager { get; set; }

    public AuthenticatedBase() : base()
    {
    }

    protected override void OnInitialized()
    {
        base.OnInitialized();
        
        // AuthenticationStateProvider.AuthenticationStateChanged += AuthenticationStateProvider_OnAuthenticationStateChanged;
    }

    protected override void OnAfterRender(bool firstRender)
    {
        base.OnAfterRender(firstRender);
        
        AuthenticationStateProvider.AuthenticationStateChanged += AuthenticationStateProvider_OnAuthenticationStateChanged;
    }

    protected virtual void AuthenticationStateProvider_OnAuthenticationStateChanged(Task<AuthenticationState> task)
    {
    }

    protected virtual bool IsAuthenticated()
    {
        return AuthenticationState.User.Identity?.IsAuthenticated ?? false;
    }

    protected virtual async Task<AuthenticationState> GetAuthenticationState()
    {
        return await AuthenticationStateProvider.GetAuthenticationStateAsync();
    }
}