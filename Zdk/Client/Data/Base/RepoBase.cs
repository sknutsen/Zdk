using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.SignalR.Client;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;

namespace Zdk.Client;

public abstract class RepoBase : IAsyncDisposable
{
    [Inject]
    protected AuthenticationStateProvider AuthenticationStateProvider { get; set; }

    [Inject]
    protected NavigationManager NavigationManager { get; set; }

    protected IList<IDisposable> subscriptions = new List<IDisposable>();
    protected HubConnection HubConnection { get; set; }

    public bool IsConnected =>
        HubConnection != null &&
        HubConnection.State == HubConnectionState.Connected;

    public RepoBase(IServiceProvider services)
    {
        AuthenticationStateProvider = services.GetService<AuthenticationStateProvider>();
        NavigationManager = services.GetService<NavigationManager>();

        HubConnection = new HubConnectionBuilder()
                .WithUrl(NavigationManager.ToAbsoluteUri($"/{GetHubName()}"))
                .WithAutomaticReconnect()
                .Build();

        HubConnection.Reconnected += HubConnectionOnReconnected;
    }

    protected abstract string GetHubName();

    public async Task Init()
    {
        if (HubConnection.State == HubConnectionState.Disconnected)
        {
            subscriptions = On().Result;

            await HubConnection.StartAsync();
        }
    }

    public async Task Disconnect()
    {
        if (IsConnected)
        {
            await HubConnection.StopAsync();
        }
    }

    protected abstract Task<IList<IDisposable>> On();

    protected abstract Task HubConnectionOnReconnected(string? arg);

    public virtual async Task JoinClientGroup()
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(ClientGroupMethodNames.JoinGroup);
        }
    }

    public virtual async Task LeaveClientGroup(int groupId)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(ClientGroupMethodNames.LeaveGroup, groupId);
        }
    }

    protected async Task SendAsync(string hubMethodName, object? arg0 = null, object? arg1 = null, object? arg2 = null, 
                                                         object? arg3 = null, object? arg4 = null, object? arg5 = null, 
                                                         object? arg6 = null, object? arg7 = null, object? arg8 = null, 
                                                         object? arg9 = null)
    {
        if (IsConnected)
        {
            await HubConnection!.SendAsync(hubMethodName, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9);
        }
    }

    public virtual async ValueTask DisposeAsync()
    {
        if (HubConnection is not null)
        {
            HubConnection.Reconnected -= HubConnectionOnReconnected;
            await HubConnection.DisposeAsync();
        }

        foreach (IDisposable sub in subscriptions)
        {
            sub.Dispose();
        }
    }
}
