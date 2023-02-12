using Microsoft.AspNetCore.Components;

namespace Zdk.Client
{
    public partial class TableComponent<TItem> : ZComponentBase
    {
        [Parameter]
        public RenderFragment TableHeader { get; set; }

        [Parameter]
        public RenderFragment<TItem> RowTemplate { get; set; }

        [Parameter]
        public RenderFragment TableFooter { get; set; }

        [Parameter]
        public IList<TItem> Items { get; set; }
    }
}