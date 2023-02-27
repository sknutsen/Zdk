using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.Helpers;

public static class DateFunc
{
    public static string ToDateText(this DateTime date, bool time = false)
    {
        string result = $"{date.DayOfWeek} {date.Day}.{date.Month}.{date.Year}";

        if (time)
        {
            result += $" {date.Hour}.{date.Minute}";
        }

        return result;
    }
}
