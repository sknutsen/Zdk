using System.Collections.Generic;

namespace Zdk.Shared.Constants;

public class BooleanDefs
{
    public static IDictionary<string, bool> List = new Dictionary<string, bool>
        {
            { "No", false },
            { "Yes", true },
        };
}

public class ItemOrderDefs
{
    public const int None = 0;
    public const int NoRepeatedCategories = 1;
    public const int OnlyRepeatCategories = 2;

    public static IDictionary<string, int> List = new Dictionary<string, int>
        {
            { "None", None },
            { "Don't repeat categories in succession", NoRepeatedCategories },
            { "One category at a time", OnlyRepeatCategories },
        };
}
