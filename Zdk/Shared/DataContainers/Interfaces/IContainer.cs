using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Zdk.Shared.Models;

namespace Zdk.Shared.DataContainers;

public interface IContainer
{
    public IEntityClass ToEntityClass();
    public void Fill(IEntityClass entity);
}
