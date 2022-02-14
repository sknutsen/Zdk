using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Zdk.Shared.Models;

namespace Zdk.Shared.Helpers;

public static class ShoppingListHelpers
{
    public static void SortAll(this ShoppingList shoppingList)
    {
        shoppingList.Items = shoppingList.Items.OrderBy(e => e.ItemId).ToList();
    }

    public static void SortAll(this IList<ShoppingList> shoppingLists)
    {
        shoppingLists = shoppingLists.OrderBy(e => e.ShoppingListId).ToList();

        foreach (ShoppingList sl in shoppingLists)
        {
            sl.SortAll();
        }
    }
}
