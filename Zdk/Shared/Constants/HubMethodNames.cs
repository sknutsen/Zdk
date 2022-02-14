namespace Zdk.Shared.Constants;

public class ClientGroupMethodNames
{
    public const string JoinGroup = "JoinGroup";
    public const string LeaveGroup = "LeaveGroup";
}

public class MainHubMethodNames
{
    public const string ReceiveMessage = "ReceiveMessage";
    public const string SendMessage = "SendMessage";

    public const string ReceiveData = "ReceiveData";
    public const string SendData = "SendData";
}

public class ShoppingListsHubMethodNames
{
    public const string ReceiveData = "ReceiveData";
    public const string ListShoppingLists = "ListShoppingLists";

    public const string GetGroups = "GetGroups";
    public const string ReceiveGroups = "ReceiveGroups";

    public const string SendShoppingList = "SendShoppingList";
    public const string SendItem = "SendItem";

    public const string UpdateShoppingList = "UpdateShoppingList";
    public const string UpdateItem = "UpdateItem";

    public const string DeleteShoppingList = "DeleteShoppingList";
    public const string DeleteItem = "DeleteItem";
}

public class UserManagementHubMethodNames
{
    public const string GetGroupMemberships = "GetGroupMemberships";

    public const string GetGroups   = "GetGroups";
    public const string NewGroup    = "NewGroup";
    public const string DeleteGroup = "DeleteGroup";
    public const string UpdateGroup = "UpdateGroup";

    public const string GetUsers    = "GetUsers";
    public const string AddUser     = "AddUser";
    public const string RemoveUser  = "RemoveUser";

    public const string GetUserSession = "GetSession";
    public const string UpdateUserSession = "UpdateSession";
}
