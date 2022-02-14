using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Zdk.Shared;
using Zdk.Shared.Constants;

namespace Zdk.Server.Hubs;

[Authorize]
public class MainHub : Hub
{
    private static readonly List<WeatherForecast> forecasts = new();

    public async Task SendMessage(string message)
    {
        await Clients.All.SendAsync(MainHubMethodNames.ReceiveMessage, this.Context.User?.Identity?.Name ?? "???", message);
    }

    public async Task SendData(WeatherForecast forecast)
    {
        forecasts.Add(forecast);

        await Clients.All.SendAsync(MainHubMethodNames.ReceiveData, forecasts);
    }

    public override async Task OnConnectedAsync()
    {
        await Clients.Caller.SendAsync(MainHubMethodNames.ReceiveData, forecasts);

        await base.OnConnectedAsync();
    }
}
