using System.Net;
using System.Net.Http;
using System.Net.Http.Json;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.Components.Forms;
using Microsoft.AspNetCore.Components.Routing;
using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.Web.Virtualization;
using Microsoft.AspNetCore.Components.WebAssembly.Authentication;
using Microsoft.AspNetCore.Components.WebAssembly.Http;
using Microsoft.AspNetCore.SignalR.Client;
using Microsoft.JSInterop;
using Zdk.Client;
using Zdk.Shared;
using Zdk.Shared.Constants;
using Zdk.Shared.Models;


namespace Zdk.Client;

public partial class Index : AuthenticatedBase, IAsyncDisposable
{
    private HubConnection? hubConnection;
    private List<string> messages = new List<string>();
    private List<WeatherForecast> data = new List<WeatherForecast>();
    private string? messageInput;
    private int tempInput;
    private string? summaryInput;

    public bool IsConnected =>
        hubConnection?.State == HubConnectionState.Connected;

    protected override async Task OnInitializedAsync()
    {
        await base.OnInitializedAsync();
    }

    protected override async Task OnParametersSetAsync()
    {
        await base.OnParametersSetAsync();

        if (IsAuthenticated())
        {
            try
            {
                hubConnection = new HubConnectionBuilder()
                    .WithUrl(NavigationManager.ToAbsoluteUri("/mainhub"))
                    .WithAutomaticReconnect()
                    .Build();

                hubConnection.On<string, string>(MainHubMethodNames.ReceiveMessage, (user, message) =>
                {
                    var encodedMsg = $"{user}: {message}";
                    messages.Add(encodedMsg);
                    StateHasChanged();
                });

                hubConnection.On<List<WeatherForecast>>(MainHubMethodNames.ReceiveData, (list) =>
                {
                    data = list;
                    StateHasChanged();
                });

                await hubConnection.StartAsync();
            }
            catch (AccessTokenNotAvailableException exception)
            {
                exception.Redirect();
            }
            catch (HttpRequestException exception)
            {
                if (exception.StatusCode == HttpStatusCode.Unauthorized)
                {
                    NavigationManager.NavigateTo($"authentication/login?returnUrl={Uri.EscapeDataString(NavigationManager.Uri)}");
                }
            }
        }
    }

    private async Task Send()
    {
        if (hubConnection is not null && IsAuthenticated())
        {
            await hubConnection.SendAsync(MainHubMethodNames.SendMessage, messageInput);
        }
    }

    private async Task SendData()
    {
        if (hubConnection is not null && IsAuthenticated())
        {
            await hubConnection.SendAsync(MainHubMethodNames.SendData, new WeatherForecast
            {
                Date = DateTime.Now,
                TemperatureC = tempInput,
                Summary = summaryInput
            });
        }
    }

    public async ValueTask DisposeAsync()
    {
        if (hubConnection is not null)
        {
            await hubConnection.DisposeAsync();
        }
    }
}
