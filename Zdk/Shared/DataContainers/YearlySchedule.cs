using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Zdk.Shared.DataContainers;

public class YearlySchedule
{
    public int Year { get; set; }
    public IEnumerable<MonthlySchedule> ScheduledItems { get; set; }
}
