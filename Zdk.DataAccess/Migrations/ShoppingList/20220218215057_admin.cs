using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Zdk.DataAccess.Migrations.ShoppingList
{
    public partial class admin : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<bool>(
                name: "is_admin",
                table: "user_sessions",
                type: "boolean",
                nullable: false,
                defaultValue: false);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "is_admin",
                table: "user_sessions");
        }
    }
}
