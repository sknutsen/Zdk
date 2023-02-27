using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Zdk.Shared.DataContainers;
using Zdk.Shared.Models;

namespace Zdk.Shared.Logic;

public static class ItemFunc
{
    public static Item GetItem(IEnumerable<ItemContainer> items, IEnumerable<ScheduledItemContainer> scheduledItems, bool allowRepeatedCategories)
    {
        Random random = new Random();

        Item result = null;

        if (allowRepeatedCategories)
        {
            result = (Item)items.ToList()[random.Next(items.Count() - 1)].ToEntityClass();
        }
        else
        {
            var list = items.ToList();
            list.Shuffle();

            foreach (var item in list)
            {
                if (!scheduledItems.First().Item.ItemCategories.Any(e => item.ItemCategories.Any(f => e.CategoryId == f.CategoryId)))
                {
                    result = (Item)item.ToEntityClass();
                    break;
                }
            }

            if (result == null)
            {
                result = (Item)items.ToList()[random.Next(items.Count() - 1)].ToEntityClass();
            }
        }

        return result;
    }

    public static void Shuffle<T>(this IList<T> list)
    {
        Random random = new Random();

        int n = list.Count;
        while (n > 1)
        {
            n--;
            int k = random.Next(n + 1);
            T value = list[k];
            list[k] = list[n];
            list[n] = value;
        }
    }
}
