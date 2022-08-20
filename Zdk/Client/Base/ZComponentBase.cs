using Microsoft.AspNetCore.Components;

namespace Zdk.Client;

public abstract class ZComponentBase : ComponentBase
{
    protected override void OnInitialized()
    {
        base.OnInitialized();
    }

    protected override async Task OnInitializedAsync()
    {
        await base.OnInitializedAsync();
    }

    public override async Task SetParametersAsync(ParameterView parameters)
    {
        await base.SetParametersAsync(parameters);
    }

    protected override void OnParametersSet()
    {
        base.OnParametersSet();
    }

    protected override async Task OnParametersSetAsync()
    {
        await base.OnParametersSetAsync();
    }

    protected override void OnAfterRender(bool firstRender)
    {
        base.OnAfterRender(firstRender);

        if (firstRender)
        {
            OnFirstRender();
        }
    }

    protected override async Task OnAfterRenderAsync(bool firstRender)
    {
        await base.OnAfterRenderAsync(firstRender);

        if (firstRender)
        {
            await OnFirstRenderAsync();
        }
    }

    protected virtual void OnFirstRender()
    {
        
    }

    protected virtual async Task OnFirstRenderAsync()
    {
        await Task.CompletedTask;
    }
}